package http

import (
	coverRoutes "coverCraft/domain/interfaces/http"
	"coverCraft/entities"
	"coverCraft/mocks"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListAllJobs(t *testing.T) {
	// Create a new Echo instance for testing
	e := echo.New()

	// Create a MockController
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock repository
	mockRepo := mocks.NewMockJobRepository(ctrl)

	// Set up expected behavior for the mock repository
	expectedJobs := []entities.Job{
		// Define expected job data for testing
		{
			ID:      uuid.New(),
			Company: "Tamara",
			Qualifications: "You have 5+ years of experience in building scalable and reliable large-scale backend " +
				"API's\nYou are very proficient in Golang\nExperience with Java or PHP\nYou have experience of Elasticsearch " +
				"and indexing products\nYou have strong experience in Architecture/System Design\nGood understanding of " +
				"SOLID principles, Design Patterns\nKnowledge of AWS, OCI, Kubernetes/Docker\nExperience with distributed " +
				"systems with messaging tools like Kafka or RabbitMQ\nExperience with Agile/Scrum\nKnowledge of Javascript, " +
				"Vue.js, HTML and CSS - is a plus\nYou are proactive, goal-orientated, " +
				"and self-structured in your approach\nYou have experience as a developer in a startup company - is a plus\nTeam player, " +
				"creative thinker, and passionate about exploring new technologies\nGood written and spoken communication skills in English\nBSc in Computer Science or equivalent",
			Description: "We firmly believe in granting individuals the freedom and accountability to excel in their areas of expertise. " +
				"As a vital member of our dynamic Engineering team, you will play a crucial role in building cutting-edge platforms that enable seamless and scalable transactions within the MENA region. Drawing on your proficiency in payment, shopping, and banking systems and algorithms, you will be responsible for designing, developing, and optimizing our secure and robust payment processing infrastructure.\n\n\nAt Tamara, we have backend engineers in almost every team, and in that role, you'll be making some of the most significant decisions for the company. You will work with our talented engineers to design, develop, and maintain complex web applications and services for our BNPL service platform.\n\n\nBesides Go, we run services using Java, PHP, and Python. In this exciting role, you will collaborate with a diverse ecosystem of talented engineers, designers, analysts, product experts, and other cross-functional team members to leverage system design, create technical solutions," +
				" dive deep into complex problems, and ultimately revolutionize the way millions of users shop, pay, and bank.",
			Location: "Saudi Arabia",
		},
	}
	mockRepo.EXPECT().ListAllJobs().Return(expectedJobs, nil)

	// Create a new HTTP request with a GET method to the ListAllJobs endpoint
	req := httptest.NewRequest(http.MethodGet, "/jobs", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Call the ListAllJobs function
	if err := coverRoutes.ListAllJobs(c, mockRepo); err != nil {
		t.Errorf("ListAllJobs returned an error: %v", err)
	}

	// Check the HTTP status code (200 OK is expected)
	if rec.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, rec.Code)
	}
}
