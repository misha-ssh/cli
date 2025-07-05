package selection

type Selector interface {
	Select(aliases []string) (string, error)
}
