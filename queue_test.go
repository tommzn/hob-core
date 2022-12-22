package core

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type QueueTestSuite struct {
	suite.Suite
}

func TestQueueTestSuite(t *testing.T) {
	suite.Run(t, new(QueueTestSuite))
}

func (suite *QueueTestSuite) TestSerializeAndDeserialize() {

	event := &GenerateReportRequest{
		Format:      ReportFormat_EXCEL,
		Type:        ReportType_MONTHLY_REPORT,
		Year:        2022,
		Month:       1,
		NamePattern: "FileName",
		Delivery: &ReportDelivery{
			S3: &S3Target{
				Region: "eu-central-5",
				Bucket: "bucket",
				Path:   "/base_path/",
			},
		},
	}

	data, err := SerializeEvent(event)
	suite.Nil(err)
	suite.True(len(data) > 10)

	event2 := &GenerateReportRequest{}
	suite.Nil(DeserializeEvent(data, event2))

	suite.Equal(event.Format, event2.Format)
	suite.Equal(event.Type, event2.Type)
	suite.Equal(event.Year, event2.Year)
	suite.Equal(event.Month, event2.Month)
	suite.Equal(event.NamePattern, event2.NamePattern)
	suite.Equal(event.Delivery.S3.Region, event2.Delivery.S3.Region)
	suite.Equal(event.Delivery.S3.Bucket, event2.Delivery.S3.Bucket)
	suite.Equal(event.Delivery.S3.Path, event2.Delivery.S3.Path)
	suite.Equal(event.Delivery.Mail, event2.Delivery.Mail)
}
