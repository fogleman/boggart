package mqtt

type ComponentSensor struct {
	*ComponentBase

	UnitOfMeasurement string `json:"unit_of_measurement"`
}

func NewComponentSensor(id string) *ComponentSensor {
	return &ComponentSensor{
		ComponentBase: NewComponentBase(id, ComponentTypeSensor),
	}
}

func (c *ComponentSensor) GetState() interface{} {
	s := c.ComponentBase.GetState().(string)
	if c.UnitOfMeasurement != "" {
		s += " " + c.UnitOfMeasurement
	}

	return s
}