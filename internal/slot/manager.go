package slot

import (
	"fmt"
	"slot-server/internal/slot/model"
	"slot-server/internal/slot/module"
	"slot-server/internal/slot/slot00"
)

type Manager struct {
	slots map[int]module.ISlot
}

func Initialize() *Manager {
	slots := make(map[int]module.ISlot)

	manager := Manager{
		slots: slots,
	}

	slots[0] = slot00.Init()

	return &manager
}

func (m *Manager) Spin(req model.SpinInput) (*model.SpinOutput, error) {

	slot := m.slots[req.Id]
	if slot == nil {
		return nil, model.Error{
			Code:    model.INVALID_SLOT_ID,
			Message: fmt.Sprintf("Invalid Slot Id : %d", req.Id),
		}
	}

	if res, err := slot.Spin(req.PrevState, req.BetCash); err != nil {
		return nil, err
	} else {
		return res, nil
	}
}
