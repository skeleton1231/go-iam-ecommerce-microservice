// Copyright 2020 Talhuang<talhuang1231@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// Package cmd create a root cobra command and add subcommands to it.
package cmd

import (
	"flag"
	"io"
	"os"

	cliflag "github.com/marmotedu/component-base/pkg/cli/flag"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/skeleton1231/go-iam-ecommerce-microservice/internal/iamctl/cmd/color"
	"github.com/skeleton1231/go-iam-ecommerce-microservice/internal/iamctl/cmd/completion"
	"github.com/skeleton1231/go-iam-ecommerce-microservice/internal/iamctl/cmd/info"
	"github.com/skeleton1231/go-iam-ecommerce-microservice/internal/iamctl/cmd/jwt"
	"github.com/skeleton1231/go-iam-ecommerce-microservice/internal/iamctl/cmd/new"
	"github.com/skeleton1231/go-iam-ecommerce-microservice/internal/iamctl/cmd/options"
	"github.com/skeleton1231/go-iam-ecommerce-microservice/internal/iamctl/cmd/policy"
	"github.com/skeleton1231/go-iam-ecommerce-microservice/internal/iamctl/cmd/secret"
	"github.com/skeleton1231/go-iam-ecommerce-microservice/internal/iamctl/cmd/set"
	"github.com/skeleton1231/go-iam-ecommerce-microservice/internal/iamctl/cmd/user"
	cmdutil "github.com/skeleton1231/go-iam-ecommerce-microservice/internal/iamctl/cmd/util"
	"github.com/skeleton1231/go-iam-ecommerce-microservice/internal/iamctl/cmd/validate"
	"github.com/skeleton1231/go-iam-ecommerce-microservice/internal/iamctl/cmd/version"
	"github.com/skeleton1231/go-iam-ecommerce-microservice/internal/iamctl/util/templates"
	genericapiserver "github.com/skeleton1231/go-iam-ecommerce-microservice/internal/pkg/server"
	"github.com/skeleton1231/go-iam-ecommerce-microservice/pkg/cli/genericclioptions"
)

// NewDefaultIAMCtlCommand creates the `iamctl` command with default arguments.
func NewDefaultIAMCtlCommand() *cobra.Command {
	return NewIAMCtlCommand(os.Stdin, os.Stdout, os.Stderr)
}

// NewIAMCtlCommand returns new initialized instance of 'iamctl' root command.
func NewIAMCtlCommand(in io.Reader, out, err io.Writer) *cobra.Command {
	// Parent command to which all subcommands are added.
	cmds := &cobra.Command{
		Use:   "iamctl",
		Short: "iamctl controls the iam platform",
		Long: templates.LongDesc(`
		iamctl controls the iam platform, is the client side tool for iam platform.

		Find more information at:
			https://github.com/marmotedu/iam/blob/master/docs/guide/en-US/cmd/iamctl/iamctl.md`),
		Run: runHelp,
		// Hook before and after Run initialize and write profiles to disk,
		// respectively.
		PersistentPreRunE: func(*cobra.Command, []string) error {
			return initProfiling()
		},
		PersistentPostRunE: func(*cobra.Command, []string) error {
			return flushProfiling()
		},
	}

	flags := cmds.PersistentFlags()
	flags.SetNormalizeFunc(cliflag.WarnWordSepNormalizeFunc) // Warn for "_" flags

	// Normalize all flags that are coming from other packages or pre-configurations
	// a.k.a. change all "_" to "-". e.g. glog package
	flags.SetNormalizeFunc(cliflag.WordSepNormalizeFunc)

	addProfilingFlags(flags)

	iamConfigFlags := genericclioptions.NewConfigFlags(true).WithDeprecatedPasswordFlag().WithDeprecatedSecretFlag()
	iamConfigFlags.AddFlags(flags)
	matchVersionIAMConfigFlags := cmdutil.NewMatchVersionFlags(iamConfigFlags)
	matchVersionIAMConfigFlags.AddFlags(cmds.PersistentFlags())

	_ = viper.BindPFlags(cmds.PersistentFlags())
	cobra.OnInitialize(func() {
		genericapiserver.LoadConfig(viper.GetString(genericclioptions.FlagIAMConfig), "iamctl")
	})
	cmds.PersistentFlags().AddGoFlagSet(flag.CommandLine)

	f := cmdutil.NewFactory(matchVersionIAMConfigFlags)

	// From this point and forward we get warnings on flags that contain "_" separators
	cmds.SetGlobalNormalizationFunc(cliflag.WarnWordSepNormalizeFunc)

	ioStreams := genericclioptions.IOStreams{In: in, Out: out, ErrOut: err}

	groups := templates.CommandGroups{
		{
			Message: "Basic Commands:",
			Commands: []*cobra.Command{
				info.NewCmdInfo(f, ioStreams),
				color.NewCmdColor(f, ioStreams),
				new.NewCmdNew(f, ioStreams),
				jwt.NewCmdJWT(f, ioStreams),
			},
		},
		{
			Message: "Identity and Access Management Commands:",
			Commands: []*cobra.Command{
				user.NewCmdUser(f, ioStreams),
				secret.NewCmdSecret(f, ioStreams),
				policy.NewCmdPolicy(f, ioStreams),
			},
		},
		{
			Message: "Troubleshooting and Debugging Commands:",
			Commands: []*cobra.Command{
				validate.NewCmdValidate(f, ioStreams),
			},
		},
		{
			Message: "Settings Commands:",
			Commands: []*cobra.Command{
				set.NewCmdSet(f, ioStreams),
				completion.NewCmdCompletion(ioStreams.Out, ""),
			},
		},
	}
	groups.Add(cmds)

	filters := []string{"options"}
	templates.ActsAsRootCommand(cmds, filters, groups...)

	cmds.AddCommand(version.NewCmdVersion(f, ioStreams))
	cmds.AddCommand(options.NewCmdOptions(ioStreams.Out))

	return cmds
}

func runHelp(cmd *cobra.Command, args []string) {
	_ = cmd.Help()
}
