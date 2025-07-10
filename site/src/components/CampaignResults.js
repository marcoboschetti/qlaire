import React from 'react';
import { 
  Copy, 
  Palette, 
  User, 
  PieChart, 
  Settings, 
  Lightbulb,
  Download,
  Share2
} from 'lucide-react';

const CampaignResults = ({ job }) => {
  if (!job || !job.ads_campaign_result) {
    return (
      <div className="results-section">
        <h2>
          <i className="fas fa-chart-line"></i> Campaign Results
        </h2>
        <div className="results-content">
          <p style={{ textAlign: 'center', color: '#6b7280', marginTop: '50px' }}>
            <i className="fas fa-lightbulb" style={{ fontSize: '3rem', marginBottom: '20px', display: 'block' }}></i>
            Fill out the form and click "Generate Campaign" to get started
          </p>
        </div>
      </div>
    );
  }

  const campaign = job.ads_campaign_result;

  const exportCampaign = () => {
    const campaignData = {
      job_id: job.id,
      generated_at: new Date().toISOString(),
      campaign: campaign
    };
    
    const blob = new Blob([JSON.stringify(campaignData, null, 2)], { type: 'application/json' });
    const url = URL.createObjectURL(blob);
    const a = document.createElement('a');
    a.href = url;
    a.download = `campaign-${job.id}.json`;
    document.body.appendChild(a);
    a.click();
    document.body.removeChild(a);
    URL.revokeObjectURL(url);
  };

  const copyToClipboard = (text) => {
    navigator.clipboard.writeText(text);
  };

  return (
    <div className="results-section">
      <div className="results-header">
        <h2>
          <i className="fas fa-chart-line"></i> Campaign Results
        </h2>
        <div className="results-actions">
          <button onClick={exportCampaign} className="btn-secondary">
            <Download size={16} />
            Export
          </button>
          <button onClick={() => copyToClipboard(JSON.stringify(campaign, null, 2))} className="btn-secondary">
            <Share2 size={16} />
            Copy JSON
          </button>
        </div>
      </div>

      <div className="campaign-result">
        {/* Ad Copy Section */}
        <div className="campaign-section">
          <h3>
            <Copy size={20} />
            Ad Copy
          </h3>
          <div className="ad-copy-grid">
            {campaign.ad_copy.map((ad, index) => (
              <div key={index} className="ad-copy-item">
                <div className="ad-copy-header">
                  <h4>Ad {index + 1}</h4>
                  <button 
                    onClick={() => copyToClipboard(`${ad.headline}\n\n${ad.description}`)}
                    className="copy-btn"
                    title="Copy to clipboard"
                  >
                    <Copy size={14} />
                  </button>
                </div>
                <div className="ad-copy-content">
                  <div className="headline">
                    <strong>Headline:</strong> {ad.headline}
                  </div>
                  <div className="description">
                    <strong>Description:</strong> {ad.description}
                  </div>
                </div>
              </div>
            ))}
          </div>
        </div>

        {/* Creative Concepts Section */}
        <div className="campaign-section">
          <h3>
            <Palette size={20} />
            Creative Concepts
          </h3>
          <div className="creative-concepts-grid">
            {campaign.creative_concepts.map((concept, index) => (
              <div key={index} className="creative-concept-item">
                <div className="concept-header">
                  <h4>{concept.concept_type}</h4>
                  <span className="concept-badge">{concept.concept_type}</span>
                </div>
                <div className="concept-content">
                  <p><strong>Description:</strong> {concept.description}</p>
                  <p><strong>Elements:</strong> {concept.elements}</p>
                </div>
              </div>
            ))}
          </div>
        </div>

        {/* Target Persona Section */}
        <div className="campaign-section">
          <h3>
            <User size={20} />
            Target Persona
          </h3>
          <div className="persona-grid">
            <div className="persona-item">
              <h4>Age</h4>
              <p>{campaign.persona_summary.age}</p>
            </div>
            <div className="persona-item">
              <h4>Gender</h4>
              <p>{campaign.persona_summary.gender}</p>
            </div>
            <div className="persona-item">
              <h4>Behavior</h4>
              <p>{campaign.persona_summary.behavior}</p>
            </div>
            <div className="persona-item">
              <h4>Interests</h4>
              <p>{campaign.persona_summary.interests}</p>
            </div>
          </div>
        </div>

        {/* Segmentation Section */}
        <div className="campaign-section">
          <h3>
            <PieChart size={20} />
            Segmentation
          </h3>
          <div className="segmentation-grid">
            <div className="segmentation-item">
              <h4>Age</h4>
              <p>{campaign.segmentation.age}</p>
            </div>
            <div className="segmentation-item">
              <h4>Gender</h4>
              <p>{campaign.segmentation.gender}</p>
            </div>
            <div className="segmentation-item">
              <h4>Behavior</h4>
              <p>{campaign.segmentation.behavior}</p>
            </div>
            <div className="segmentation-item">
              <h4>Devices</h4>
              <p>{campaign.segmentation.devices}</p>
            </div>
            <div className="segmentation-item">
              <h4>Interests</h4>
              <p>{campaign.segmentation.interests}</p>
            </div>
            <div className="segmentation-item">
              <h4>Location</h4>
              <p>{campaign.segmentation.location}</p>
            </div>
          </div>
        </div>

        {/* Campaign Configuration Section */}
        <div className="campaign-section">
          <h3>
            <Settings size={20} />
            Campaign Configuration
          </h3>
          <div className="config-grid">
            <div className="config-item">
              <h4>Objective</h4>
              <p>{campaign.campaign_config.objective}</p>
            </div>
            <div className="config-item">
              <h4>Placements</h4>
              <p>{campaign.campaign_config.placements}</p>
            </div>
            <div className="config-item">
              <h4>Budget</h4>
              <p>{campaign.campaign_config.budget}</p>
            </div>
          </div>
          
          <div className="ab-testing-section">
            <h4>A/B Testing Variants</h4>
            <div className="ab-testing-grid">
              {campaign.campaign_config.a_b_testing.map((test, index) => (
                <div key={index} className="ab-test-item">
                  <h5>{test.test_name}</h5>
                  <p>{test.variants}</p>
                </div>
              ))}
            </div>
          </div>
        </div>

        {/* Key Insights Section */}
        <div className="campaign-section">
          <h3>
            <Lightbulb size={20} />
            Key Insights
          </h3>
          <ul className="insights-list">
            {campaign.key_insights.map((insight, index) => (
              <li key={index}>
                <i className="fas fa-check"></i> {insight}
              </li>
            ))}
          </ul>
        </div>
      </div>
    </div>
  );
};

export default CampaignResults; 