package build

import (
	// "fmt"

	"fmt"

	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/cli-runtime/pkg/genericiooptions"
	cmdutil "k8s.io/kubectl/pkg/cmd/util"
	"k8s.io/kubectl/pkg/util/i18n"
)

type BuildOptions struct {
	Namespace         string
	NoHeaders         bool
	ExplicitNamespace bool
	AllNamespaces     bool
	LabelSelector     string
	FieldSelector     string
	ChunkSize         int64
	genericiooptions.IOStreams
}

func NewBuildCmd(f cmdutil.Factory, ioStreams genericiooptions.IOStreams) *cobra.Command {
	buildIt := BuildOptions{
		IOStreams: ioStreams,
		ChunkSize: cmdutil.DefaultChunkSize,
	}
	cmd := &cobra.Command{
		Use:                   "build it",
		DisableFlagsInUseLine: true,
		Short:                 i18n.T("Apply a configuration to a resource by file name or stdin"),
		Run: func(cmd *cobra.Command, args []string) {
			var value = []string{"pods"}
			cmdutil.CheckErr(buildIt.Run(f, cmd, value))
		},
	}
	cmd.Flags().BoolVarP(&buildIt.AllNamespaces, "all-namespaces", "A", buildIt.AllNamespaces, "If present, list the requested object(s) across all namespaces. Namespace in current context is ignored even if specified with --namespace.")
	return cmd
}

func (b *BuildOptions) Run(f cmdutil.Factory, cmd *cobra.Command, args []string) error {
	var err error
	// get the namespace provided
	b.Namespace, b.ExplicitNamespace, err = f.ToRawKubeConfigLoader().Namespace()
	if err != nil {
		return err
	}
	if b.AllNamespaces {
		b.ExplicitNamespace = false
	}

	r := f.NewBuilder().
		Unstructured().
		NamespaceParam(b.Namespace).DefaultNamespace().AllNamespaces(b.AllNamespaces).
		LabelSelectorParam(b.LabelSelector).
		FieldSelectorParam(b.FieldSelector).
		RequestChunksOf(b.ChunkSize).
		ResourceTypeOrNameArgs(true, args...).
		ContinueOnError().
		Latest().
		Flatten().
		Do()

	infos, err := r.Infos()
	objs := make([]runtime.Object, len(infos))
	for ix := range infos {
		objs[ix] = infos[ix].Object
		fmt.Println(objs[ix])
		fmt.Println("x", 5)
	}
	return err
}
