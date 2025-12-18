package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/opencontainers/umoci"
	"github.com/opencontainers/umoci/oci/cas/dir"
	"github.com/opencontainers/umoci/oci/casext"
	"github.com/opencontainers/umoci/oci/layer"

	rspec "github.com/opencontainers/runtime-spec/specs-go"

	"github.com/spf13/cobra"
)

type unpackOptions struct {
	global *globalOptions
}

func unpackCmd(global *globalOptions) *cobra.Command {
	opts := unpackOptions{global: global}
	cmd := &cobra.Command{
		Use:     "unpack OCI-DIR DESTINATION-ROOTFS",
		Short:   "Unpack an OCI-DIR from one location to another",
		RunE:    commandAction(opts.run),
		Example: `DockRoot unpack ./alpine-oci:latest alpine-rootfs`,
	}

	return cmd
}

func (opts *unpackOptions) run(args []string, stdout io.Writer) (retErr error) {
	log.Println("args=", args)
	if len(args) != 2 {
		return fmt.Errorf("invalid number of arguments")
	}
	imagePath := args[0]
	ss := strings.Split(imagePath, ":")
	if len(ss) != 2 {
		return fmt.Errorf("invalid image path")
	}
	imagePath = ss[0]
	fromName := ss[1]
	bundlePath := args[1]
	return unpack(imagePath, fromName, bundlePath)
}

func unpack(imagePath, fromName, bundlePath string) error {
	var meta umoci.Meta
	meta.Version = umoci.MetaVersion
	unpackOptions := layer.UnpackOptions{
		OnDiskFormat: layer.DirRootfs{
			MapOptions: meta.MapOptions,
		},
		KeepDirlinks: true,
	}
	// Get a reference to the CAS.
	engine, err := dir.Open(imagePath)
	if err != nil {
		return fmt.Errorf("open CAS: %w", err)
	}
	engineExt := casext.NewEngine(engine)
	defer engine.Close()
	err = umoci.Unpack(engineExt, fromName, bundlePath, unpackOptions)

	{
		// Delete useless files
		os.Remove(filepath.Join(bundlePath, "umoci.json"))
		files, err2 := filepath.Glob(filepath.Join(bundlePath, "*.mtree"))
		if err2 != nil {
			return err
		}
		for _, file := range files {
			os.Remove(file)
		}
	}

	return err
}

func getSpecConfig(bundlePath string) (*rspec.Spec, error) {
	f, err := os.Open(bundlePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	spec := &rspec.Spec{}
	err = json.NewDecoder(f).Decode(spec)
	if err != nil {
		return nil, err
	}
	return spec, nil
}

func saveSpecConfig(bundlePath string, spec *rspec.Spec) error {
	configFile, err := os.OpenFile(bundlePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return nil
	}
	enc := json.NewEncoder(configFile)
	enc.SetIndent("", "\t")
	if err := enc.Encode(spec); err != nil {
		return fmt.Errorf("write config.json: %w", err)
	}
	return nil
}
