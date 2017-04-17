package inmem

import "github.com/go-kit/kit/examples/shipping/model/voyage"

type voyageRepository struct {
	voyages map[voyage.Number]*voyage.Voyage
}

func (r *voyageRepository) Find(voyageNumber voyage.Number) (*voyage.Voyage, error) {
	if v, ok := r.voyages[voyageNumber]; ok {
		return v, nil
	}

	return nil, voyage.ErrUnknown
}

// NewVoyageRepository returns a new instance of a in-memory voyage repository.
func NewVoyageRepository() voyage.Repository {
	r := &voyageRepository{
		voyages: make(map[voyage.Number]*voyage.Voyage),
	}

	r.voyages[voyage.V100.Number] = voyage.V100
	r.voyages[voyage.V300.Number] = voyage.V300
	r.voyages[voyage.V400.Number] = voyage.V400

	r.voyages[voyage.V0100S.Number] = voyage.V0100S
	r.voyages[voyage.V0200T.Number] = voyage.V0200T
	r.voyages[voyage.V0300A.Number] = voyage.V0300A
	r.voyages[voyage.V0301S.Number] = voyage.V0301S
	r.voyages[voyage.V0400S.Number] = voyage.V0400S

	return r
}