package cmd

import (
	"io"
	"kubing/pkg/cmd/build"
	"os"

	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/cli-runtime/pkg/genericiooptions"

	cliflag "k8s.io/component-base/cli/flag"
	"k8s.io/kubectl/pkg/cmd/options"
	cmdutil "k8s.io/kubectl/pkg/cmd/util"
	"k8s.io/kubectl/pkg/util/templates"
)

type commandFlags struct {
	ConfigFlags   *genericclioptions.ConfigFlags
	LabelSelector string
	FieldSelector string
	genericclioptions.IOStreams
}

func defaultConfigFlags() *genericclioptions.ConfigFlags {
	return genericclioptions.NewConfigFlags(true).WithDeprecatedPasswordFlag().WithDiscoveryBurst(300).WithDiscoveryQPS(50.0)
}

func NewKubingCommand() *cobra.Command {

	ioStreams := genericiooptions.IOStreams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr}
	invoke := commandFlags{
		ConfigFlags: defaultConfigFlags().WithWarningPrinter(ioStreams),
		IOStreams:   ioStreams,
	}

	// cmdutil.AddLabelSelectorFlagVar(cmds, &invoke.LabelSelector)
	return NewKubectlCommand(invoke)
}

func NewKubectlCommand(invoke commandFlags) *cobra.Command {
	cmds := &cobra.Command{
		Use:   "kubings",
		Short: "kubing all the cmds",
		Long:  "lets do the kubing thing in the command list",
		Run:   runHelp,
	}
	cmds.SetGlobalNormalizationFunc(cliflag.WarnWordSepNormalizeFunc)
	flags := cmds.PersistentFlags()
	kubeConfigFlags := invoke.ConfigFlags
	if kubeConfigFlags == nil {
		kubeConfigFlags = defaultConfigFlags().WithWarningPrinter(invoke.IOStreams)
	}
	kubeConfigFlags.AddFlags(flags)
	f := cmdutil.NewFactory(invoke.ConfigFlags)
	filters := []string{"options"}
	var group templates.CommandGroups
	templates.ActsAsRootCommand(cmds, filters, group...)
	cmds.AddCommand(options.NewCmdOptions(invoke.IOStreams.Out))
	cmds.AddCommand(build.NewBuildCmd(f, invoke.IOStreams))
	return cmds
}

func runHelp(cmd *cobra.Command, args []string) {
	var out io.Writer
	cmd.Usage()
	cmd.SetOut(out)
	cmd.SetErr(out)

}
