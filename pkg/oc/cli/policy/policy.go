package policy

import (
	"github.com/spf13/cobra"
	kcmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"
	"k8s.io/kubernetes/pkg/kubectl/genericclioptions"

	adminpolicy "github.com/openshift/origin/pkg/oc/cli/admin/policy"
	"github.com/openshift/origin/pkg/oc/cli/policy/cani"
)

const PolicyRecommendedName = "policy"

func NewCmdPolicy(name, fullName string, f kcmdutil.Factory, streams genericclioptions.IOStreams) *cobra.Command {
	// Parent command to which all subcommands are added.
	cmds := &cobra.Command{
		Use:   name,
		Short: "Manage authorization policy",
		Long:  `Manage authorization policy`,
		Run:   kcmdutil.DefaultSubCommandRun(streams.ErrOut),
	}

	cmds.AddCommand(adminpolicy.NewCmdWhoCan(adminpolicy.WhoCanRecommendedName, fullName+" "+adminpolicy.WhoCanRecommendedName, f, streams.Out))
	cmds.AddCommand(cani.NewCmdCanI(cani.CanIRecommendedName, fullName+" "+cani.CanIRecommendedName, f, streams.Out))
	cmds.AddCommand(adminpolicy.NewCmdSccSubjectReview(adminpolicy.SubjectReviewRecommendedName, fullName+" "+adminpolicy.SubjectReviewRecommendedName, f, streams.Out))
	cmds.AddCommand(adminpolicy.NewCmdSccReview(adminpolicy.ReviewRecommendedName, fullName+" "+adminpolicy.ReviewRecommendedName, f, streams.Out))

	cmds.AddCommand(adminpolicy.NewCmdAddRoleToUser(adminpolicy.AddRoleToUserRecommendedName, fullName+" "+adminpolicy.AddRoleToUserRecommendedName, f, streams.Out, streams.ErrOut))
	cmds.AddCommand(adminpolicy.NewCmdRemoveRoleFromUser(adminpolicy.RemoveRoleFromUserRecommendedName, fullName+" "+adminpolicy.RemoveRoleFromUserRecommendedName, f, streams.Out, streams.ErrOut))
	cmds.AddCommand(adminpolicy.NewCmdRemoveUserFromProject(adminpolicy.RemoveUserRecommendedName, fullName+" "+adminpolicy.RemoveUserRecommendedName, f, streams.Out))
	cmds.AddCommand(adminpolicy.NewCmdAddRoleToGroup(adminpolicy.AddRoleToGroupRecommendedName, fullName+" "+adminpolicy.AddRoleToGroupRecommendedName, f, streams.Out, streams.ErrOut))
	cmds.AddCommand(adminpolicy.NewCmdRemoveRoleFromGroup(adminpolicy.RemoveRoleFromGroupRecommendedName, fullName+" "+adminpolicy.RemoveRoleFromGroupRecommendedName, f, streams.Out, streams.ErrOut))
	cmds.AddCommand(adminpolicy.NewCmdRemoveGroupFromProject(adminpolicy.RemoveGroupRecommendedName, fullName+" "+adminpolicy.RemoveGroupRecommendedName, f, streams.Out))

	return cmds
}
