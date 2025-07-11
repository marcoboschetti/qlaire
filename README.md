# Qlaire - AI-Powered Ad Campaign Generator

Qlaire is a sophisticated API that generates tailored ad campaigns using LLMs and Qloo's cultural insights platform. It combines the power of AI with cultural data to create highly targeted and effective advertising campaigns.

## 🚀 Features

- **AI-Powered Campaign Generation**: Uses LLMs to create compelling ad copy and creative concepts
- **Cultural Intelligence**: Leverages Qloo's platform for cultural relevance and audience insights
- **Multi-Platform Support**: Generates campaigns for Meta Ads, Google Ads, TikTok, LinkedIn, and Twitter
- **Real-time Processing**: Asynchronous job processing with real-time status updates
- **Comprehensive Output**: Complete campaign packages including ad copy, creative concepts, segmentation, and insights
- **Beautiful Web Interface**: Modern, responsive frontend for easy campaign creation

## 🏗️ Architecture

The application follows a clean architecture pattern with clear separation of concerns:

```
src/
├── api/          # HTTP handlers and controllers
├── clients/      # External service clients (Qloo, LLM)
├── entities/     # Data models and structures
├── repository/   # Data persistence layer
├── service/      # Business logic and orchestration
└── site/         # Frontend web interface
```

## 📋 Prerequisites

- Go 1.22.5 or higher
- Qloo Hackathon API Token
- Groq API Token (for LLM processing)

## 🛠️ Setup

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd qlaire
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Configure environment variables**
   ```bash
   cp env.example .env
   ```
   
   Edit `.env` and add your API keys:
   ```env
   QLOO_HACKATHON_API_TOKEN=your_qloo_api_token_here
   GROQ_API_TOKEN=your_groq_api_token_here
   PORT=8080
   ```

4. **Run the application**
   ```bash
   go run main.go
   ```

5. **Access the web interface**
   Open your browser and navigate to `http://localhost:8080`

## 📚 API Documentation

### Endpoints

#### 1. Start Ad Campaign Job
**POST** `/v1/api/ads/insights`

Creates a new ad campaign generation job.

**Request Body:**
```json
{
  "product": "Premium Wireless Headphones",
  "target_platform": "Meta Ads",
  "title": "Summer Sale - Premium Audio Experience"
}
```

**Response:**
```json
{
  "job": {
    "id": "uuid-string",
    "status": "pending",
    "job_inputs": {
      "product": "Premium Wireless Headphones",
      "target_platform": "Meta Ads",
      "title": "Summer Sale - Premium Audio Experience"
    }
  }
}
```

#### 2. Get Job Status
**GET** `/v1/api/ads/insights/{jobId}`

Retrieves the current status and results of a job.

**Response:**
```json
{
  "job": {
    "id": "uuid-string",
    "status": "completed",
    "job_inputs": { ... },
    "generated_seed": {
      "query": "Beats by Dre",
      "type": "urn:entity:brand"
    },
    "search_results": [ ... ],
    "popularity_insights": [ ... ],
    "demographics": [ ... ],
    "ads_campaign_result": {
      "ad_copy": [
        {
          "headline": "Experience Premium Sound",
          "description": "Discover crystal-clear audio with our wireless headphones"
        }
      ],
      "creative_concepts": [
        {
          "type": "Image",
          "description": "Lifestyle shot of young professional enjoying music",
          "elements": "Headphones, urban background, warm lighting"
        }
      ],
      "persona_summary": {
        "age": "25-34",
        "gender": "All genders",
        "behavior": "Tech-savvy, music enthusiasts",
        "interests": "Technology, music, lifestyle"
      },
      "segmentation": {
        "age": "25-34 (60%), 18-24 (25%), 35-44 (15%)",
        "gender": "Male (55%), Female (45%)",
        "behavior": "Early adopters, social media active",
        "devices": "Mobile (70%), Desktop (30%)",
        "interests": "Technology, music, fitness, lifestyle",
        "location": "Urban areas, tech hubs"
      },
      "campaign_config": {
        "objective": "Brand awareness and conversions",
        "placements": "Facebook Feed, Instagram Stories, Audience Network",
        "budget": "$50-200 daily",
        "a_b_testing": [
          {
            "test_name": "Ad Copy Variations",
            "variants": "Premium vs Affordable messaging"
          }
        ]
      },
      "key_insights": [
        "Target audience shows high affinity for premium audio brands",
        "Mobile-first approach recommended for this demographic",
        "Lifestyle imagery performs better than product-focused ads"
      ]
    }
  }
}
```

### Job Status Flow

Jobs progress through the following states:

1. **`pending`** - Job created, waiting to start
2. **`generating_seed`** - LLM generating Qloo search seed
3. **`searching_entity`** - Searching for relevant entities in Qloo
4. **`fetching_insights`** - Retrieving popularity insights
5. **`fetching_demographics`** - Analyzing demographic data
6. **`generating_output`** - LLM generating final campaign
7. **`completed`** - Campaign successfully generated
8. **`failed`** - Job failed (check `final_error` field)

## 🎨 Web Interface

The application includes a modern React-based web interface that provides:

- **Intuitive Form**: Easy campaign input with validation
- **Real-time Status**: Live progress tracking with visual indicators
- **Step-by-step Progress**: Clear indication of current processing stage
- **Collapsible Step Details**: High-level summaries with expandable detailed views
- **Qloo Cultural Intelligence**: Emphasizes the value of Qloo's Taste AI™ insights
- **Prominent Campaign Results**: Full-width display of generated campaigns
- **Rich Results Display**: Beautifully formatted campaign results with export functionality
- **Mobile Responsive**: Works seamlessly on all devices
- **Copy to Clipboard**: Easy copying of ad copy and campaign data
- **Export Functionality**: Download campaign results as JSON with Qloo insights metadata

### Frontend Development

The React frontend is located in the `site/` directory. For development:

```bash
cd site
npm install
npm start
```

This will start the React development server on `http://localhost:3000` with hot reloading.

For production builds:

```bash
cd site
npm run build
./build.sh
```

See `site/README.md` for detailed frontend documentation.

## 🔧 Configuration

### Environment Variables

| Variable | Description | Required |
|----------|-------------|----------|
| `QLOO_HACKATHON_API_TOKEN` | Qloo API token for cultural insights | Yes |
| `GROQ_API_TOKEN` | Groq API token for LLM processing | Yes |
| `PORT` | Server port (default: 8080) | No |

### Supported Platforms

- Meta Ads (Facebook/Instagram)
- Google Ads
- TikTok Ads
- LinkedIn Ads
- Twitter Ads

## 🚀 Usage Examples

### Using the Web Interface

1. Open `http://localhost:8080` in your browser
2. Fill in the campaign details:
   - Product name
   - Target platform
   - Campaign title
3. Click "Generate Campaign"
4. Monitor the real-time progress
5. View the complete campaign results

### Using the API Directly

```bash
# Start a campaign
curl -X POST http://localhost:8080/v1/api/ads/insights \
  -H "Content-Type: application/json" \
  -d '{
    "product": "Premium Wireless Headphones",
    "target_platform": "Meta Ads",
    "title": "Summer Sale - Premium Audio Experience"
  }'

# Check job status
curl http://localhost:8080/v1/api/ads/insights/{job-id}
```

## 🔍 How It Works

1. **Input Processing**: User provides product details and target platform
2. **Seed Generation**: LLM analyzes the input and generates a culturally relevant search seed
3. **Entity Search**: Qloo searches for relevant entities based on the seed
4. **Insights Gathering**: Retrieves popularity and demographic data for the entities
5. **Campaign Generation**: LLM creates a comprehensive ad campaign using all gathered data
6. **Result Delivery**: Returns structured campaign data ready for deployment

## 🛡️ Error Handling

The API includes comprehensive error handling:

- **Input Validation**: Validates all required fields
- **API Error Handling**: Graceful handling of external API failures
- **Job Status Tracking**: Clear error messages for failed jobs
- **Retry Logic**: Automatic retries for transient failures

## 📊 Performance

- **Asynchronous Processing**: Jobs run in the background
- **Real-time Updates**: Status polling every 2 seconds
- **Memory Management**: In-memory job storage with cleanup
- **Scalable Architecture**: Easy to extend with persistent storage

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## 📄 License

This project is licensed under the MIT License.

## 🆘 Support

For support or questions, please open an issue in the repository or contact the development team.

---

**Qlaire** - Transforming advertising with AI and cultural intelligence. 