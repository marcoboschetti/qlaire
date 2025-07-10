# Qlaire React Frontend

This is the React frontend for the Qlaire AI-Powered Ad Campaign Generator.

## 🚀 Development

### Prerequisites

- Node.js 16+ 
- npm or yarn

### Setup

1. **Install dependencies**
   ```bash
   cd site
   npm install
   ```

2. **Start development server**
   ```bash
   npm start
   ```
   
   The React app will run on `http://localhost:3000` and proxy API requests to the Go backend on `http://localhost:8080`.

### Building for Production

1. **Build the React app**
   ```bash
   npm run build
   ```

2. **Or use the build script**
   ```bash
   ./build.sh
   ```

   This will build the app and copy the files to the main site directory for the Go server to serve.

## 📁 Project Structure

```
site/
├── public/                 # Static files
│   └── index.html         # Main HTML template
├── src/                   # React source code
│   ├── components/        # React components
│   │   ├── Header.js      # App header
│   │   ├── CampaignForm.js # Campaign input form
│   │   ├── JobStatus.js   # Job status and progress
│   │   ├── JobSteps.js    # Detailed step information
│   │   └── CampaignResults.js # Final campaign results
│   ├── App.js            # Main app component
│   ├── App.css           # App-specific styles
│   ├── index.js          # React entry point
│   └── index.css         # Global styles
├── package.json          # Dependencies and scripts
└── build.sh             # Production build script
```

## 🎨 Features

### Components

- **Header**: App title and branding
- **CampaignForm**: Input form for campaign details
- **JobStatus**: Real-time job status with progress indicators
- **JobSteps**: Detailed view of each processing step
- **CampaignResults**: Comprehensive campaign results display

### Key Features

- **Real-time Updates**: Live polling of job status
- **Step-by-step Progress**: Visual indicators for each processing stage
- **Detailed Step Information**: View intermediate data from each step
- **Export Functionality**: Download campaign results as JSON
- **Copy to Clipboard**: Easy copying of ad copy and campaign data
- **Responsive Design**: Works on all device sizes
- **Modern UI**: Clean, professional interface with smooth animations

### Job Steps Display

The app shows detailed information for each processing step:

1. **Job Inputs**: Original campaign parameters
2. **Generated Seed**: LLM-generated Qloo search seed
3. **Search Results**: Entities found by Qloo
4. **Popularity Insights**: Related entities and popularity scores
5. **Demographics**: Age and gender distribution data

### Campaign Results

The final results are displayed in organized sections:

- **Ad Copy**: Multiple ad variations with copy buttons
- **Creative Concepts**: Visual and video concepts
- **Target Persona**: Detailed audience profile
- **Segmentation**: Targeting parameters
- **Campaign Configuration**: Settings and A/B testing variants
- **Key Insights**: Strategic recommendations

## 🔧 Configuration

### Environment Variables

- `REACT_APP_API_URL`: Backend API URL (defaults to empty for proxy)

### API Integration

The frontend communicates with the Go backend through:

- `POST /v1/api/ads/insights` - Start new campaign
- `GET /v1/api/ads/insights/{jobId}` - Get job status

## 🛠️ Development Tips

### Adding New Components

1. Create the component in `src/components/`
2. Import and use in `App.js`
3. Add any component-specific styles to `App.css`

### Styling

- Global styles are in `src/index.css`
- Component-specific styles are in `src/App.css`
- Uses CSS Grid and Flexbox for responsive layouts
- Color scheme matches the original design

### State Management

- Uses React hooks for local state
- Job polling is handled in the main App component
- Error handling is centralized

## 🚀 Deployment

1. Build the React app: `npm run build`
2. The Go server will serve the built files from `./site/`
3. Ensure all static assets are properly served

## 📱 Browser Support

- Chrome 90+
- Firefox 88+
- Safari 14+
- Edge 90+

## 🤝 Contributing

1. Follow the existing code style
2. Add proper error handling
3. Test on different screen sizes
4. Update documentation as needed 