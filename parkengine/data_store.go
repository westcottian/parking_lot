package parkengine

type DataStore interface {
	Park(string, string) (StoreyResponse, error)
	LeaveByPosition(int) (StoreyResponse, error)
	FindByRegistrationNumber(string) (StoreyResponse, error)
	FindAllByColor(string, string) (StoreyResponse, error)
	All() (StoreyResponse, error)
	AddStorey(int) (StoreyResponse, error)
}
