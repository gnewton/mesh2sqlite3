package lib

type MeshTree struct {
	ID             int64
	DescriptorUI   string `sql:"size:16"`
	DescriptorName string
	Tree           string
	Year           int16
	Depth          int
	T0             *string `sql:"size:1"`
	T1             *string `sql:"size:3"`
	T2             *string `sql:"size:3"`
	T3             *string `sql:"size:3"`
	T4             *string `sql:"size:3"`
	T5             *string `sql:"size:3"`
	T6             *string `sql:"size:3"`
	T7             *string `sql:"size:3"`
	T8             *string `sql:"size:3"`
	T9             *string `sql:"size:3"`
	T10            *string `sql:"size:3"`
	T11            *string `sql:"size:3"`
	T12            *string `sql:"size:3"`
	T13            *string `sql:"size:3"`
}
