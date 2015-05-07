package influxdb

type ConsistencyLevel int

const (
	// ConsistencyLevelAny allows for hinted hand off, potentially no write happened yet
	ConsistencyLevelAny ConsistencyLevel = iota
	// ConsistencyLevelOne
	ConsistencyLevelOne
	ConsistencyLevelQuorum
	ConsistencyLevelOwner
	ConsistencyLevelAll
)

func newConsistencyPolicyN(need int) ConsistencyPolicy {
	return &policyNum{
		need: need,
	}
}

func newConsistencyOwnerPolicy(ownerID int) ConsistencyPolicy {
	return &policyOwner{
		ownerID: ownerID,
	}
}

type ConsistencyPolicy interface {
	IsDone(writerID int, err error) bool
}

// policyNum implements One, Quorum, and All
type policyNum struct {
	failed, succeeded, need int
}

func (p *policyNum) IsDone(writerID int, err error) bool {
	if err == nil {
		p.succeeded++
		return p.succeeded >= p.need
	}
	p.failed++
	return p.need-p.failed-p.succeeded >= p.need-p.succeeded

}

type policyOwner struct {
	ownerID int
}

func (p *policyOwner) IsDone(writerID int, err error) bool {
	return p.ownerID == writerID
}
