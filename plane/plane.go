package plane

type IPlane interface {
	GetTotalAdults() int
	GetTotalKids() int
}

type Passanger struct{}

type Plane struct {
	Passangers IPlane
}

func (p *Plane) GetInformationPassanger() (int, int) {
	return p.Passangers.GetTotalAdults(), p.Passangers.GetTotalKids()
}

func (p *Plane) IsPossibleToTakeOff() bool {
	totalKids := p.Passangers.GetTotalKids()
	totalAdults := p.Passangers.GetTotalAdults()

	return totalKids < totalAdults
}

func (pa *Passanger) GetTotalAdults() int {
	return 50
}

func (pa *Passanger) GetTotalKids() int {
	return 10
}
