package campaign

type CampaingResponse struct {
	Name                 string
	Status               string
	ID                   string
	Content              string
	AmountOfEmailsToSend int
	CreatedBy            string
}
