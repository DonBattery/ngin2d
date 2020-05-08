package ngin2d

type Taggable interface {
	GetTags() []string
	ClearTags()
	AddTags(...string)
	RemoveTags(...string)
	HasTags(...string) bool
}

type TaggableObject struct {
	tags []string
}

func (to *TaggableObject) GetTags() []string {
	return to.tags
}

func (to *TaggableObject) AddTags(tags ...string) {
	to.tags = append(to.tags, tags...)
}

func (to *TaggableObject) RemoveTags(tags ...string) {
	for _, tag := range tags {
		for i := len(to.tags) - 1; i >= 0; i-- {
			if tag == to.tags[i] {
				to.tags = append(to.tags[:i], to.tags[i+1:]...)
			}
		}
	}
}

func (to *TaggableObject) ClearTags() {
	to.tags = []string{}
}

func (to *TaggableObject) HasTags(tags ...string) bool {
	for _, tag := range tags {
		found := false
		for _, elemTag := range to.tags {
			if tag == elemTag {
				found = true
				continue
			}
		}
		if !found {
			return false
		}
	}
	return true
}
