package demo

func GetExperiments() []Experiment {
	return []Experiment{
		{
			Number:         1,
			Material:       "Ст_20_300",
			Responsible:    "Анатольев А. Н.",
			ExperimentType: "растяжение",
		},
		{
			Number:         2,
			Material:       "Ст_20_320",
			Responsible:    "Анатольев А. Н.",
			ExperimentType: "растяжение",
		},
		{
			Number:         3,
			Material:       "Ст_15_300",
			Responsible:    "Анатольев А. Н.",
			ExperimentType: "сжатие",
		},
	}
}

var index = 0

func GetExperimentsRandomly(count int) []Experiment {
	experiments := make([]Experiment, 0, count)
	lastIndex := index + count
	for ; index < lastIndex; index++ {
		experiments = append(experiments, Experiment{
			Number:         index + 1,
			Material:       "Ст_20_300",
			Responsible:    "Анатольев А. Н.",
			ExperimentType: "растяжение",
		})
	}
	return experiments
}

type Experiment struct {
	Number         int    `json:"number"`
	Material       string `json:"material"`
	Responsible    string `json:"responsible"`
	ExperimentType string `json:"experiment_type"`
}
