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

type Experiment struct {
	Number         int    `json:"number"`
	Material       string `json:"material"`
	Responsible    string `json:"responsible"`
	ExperimentType string `json:"experiment_type"`
}
