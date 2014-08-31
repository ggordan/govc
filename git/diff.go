package git

import git2go "github.com/libgit2/git2go"

type changeSet struct {
	Prev, Next, Patch string
}

type DiffJSON struct {
	FilesChanged []changeSet
}

func (dj *DiffJSON) Generate(commit *git2go.Commit) {
	// Try to diff
	parentTree, _ := commit.Parent(0).Tree()
	commitTree, _ := commit.Tree()

	do, _ := git2go.DefaultDiffOptions()
	diff, _ := commit.Owner().DiffTreeToTree(commitTree, parentTree, &do)

	// delta = file changed
	filesChanged, _ := diff.NumDeltas()

	for j := 0; j < filesChanged; j++ {
		delta, _ := diff.GetDelta(j)
		patch, _ := diff.Patch(j)

		patchString, _ := patch.String()
		// h := pygments.Highlight(patchString, "json", "html", "utf-8")

		dj.FilesChanged = append(dj.FilesChanged, changeSet{
			Prev:  delta.OldFile.Path,
			Next:  delta.NewFile.Path,
			Patch: patchString,
		})
	}
}
