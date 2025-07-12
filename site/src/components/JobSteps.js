import React, { useState } from 'react';
import { Search, TrendingUp, Users, FileText, Target, ChevronDown, ChevronUp, Brain, Zap } from 'lucide-react';

const JobSteps = ({ job }) => {
  const [expandedSteps, setExpandedSteps] = useState(new Set());

  const toggleStep = (stepName) => {
    const newExpanded = new Set(expandedSteps);
    if (newExpanded.has(stepName)) {
      newExpanded.delete(stepName);
    } else {
      newExpanded.add(stepName);
    }
    setExpandedSteps(newExpanded);
  };

  const renderSeedStep = () => {
    if (!job.generated_seed) return null;
    
    const isExpanded = expandedSteps.has('seed');
    
    return (
      <div className="step-summary">
        <div className="step-header" onClick={() => toggleStep('seed')}>
          <div className="step-info">
            <Search size={20} />
            <div>
              <h4>Cultural Seed Generated</h4>
              <p>We identified "{job.generated_seed.query}" as the most culturally relevant reference to query Qloo</p>
            </div>
          </div>
          {isExpanded ? <ChevronUp size={20} /> : <ChevronDown size={20} />}
        </div>
        
        {isExpanded && (
          <div className="step-content">
            <div className="data-grid">
              <div className="data-item">
                <h5>Cultural Entity</h5>
                <p>{job.generated_seed.query}</p>
              </div>
              <div className="data-item">
                <h5>Entity Type</h5>
                <p>{job.generated_seed.type}</p>
              </div>
            </div>
          </div>
        )}
      </div>
    );
  };

  const renderSearchStep = () => {
    if (!job.search_results || job.search_results.length === 0) return null;
    
    const isExpanded = expandedSteps.has('search');
    
    return (
      <div className="step-summary">
        <div className="step-header" onClick={() => toggleStep('search')}>
          <div className="step-info">
            <Search size={20} />
            <div>
              <h4>Cultural Entities Found</h4>
              <p>Qloo discovered {job.search_results.length} related cultural entities</p>
            </div>
          </div>
          {isExpanded ? <ChevronUp size={20} /> : <ChevronDown size={20} />}
        </div>
        
        {isExpanded && (
          <div className="step-content">
            <div className="qloo-highlight">
              <Brain size={16} />
              <span>Qloo's cultural graph connected your product to relevant entities</span>
            </div>
            <div className="data-grid">
              {job.search_results.map((result, index) => {
                // Find matching popularity insight for this search result
                const matchingInsight = job.insights_response?.find(insight => 
                  insight.EntityID === result.entity_id || insight.name === result.name
                );
                
                return (
                  <div key={index} className="data-item">
                    <h5>{result.name}</h5>
                    <p><strong>Type:</strong> {result.types.join(', ')}</p>
                    {result.short_desc && (
                      <p><strong>Description:</strong> {result.short_desc}</p>
                    )}
                    {matchingInsight && (
                      <p><strong>Popularity:</strong> {matchingInsight.popularity.toFixed(2)}</p>
                    )}
                  </div>
                );
              })}
            </div>
          </div>
        )}
      </div>
    );
  };

  const renderInsightsStep = () => {
    if (!job.popularity_insights || job.popularity_insights.length === 0) return null;
    
    const isExpanded = expandedSteps.has('insights');
    const topInsights = job.popularity_insights
      .sort((a, b) => b.popularity - a.popularity)
      .slice(0, 3);
    
    return (
      <div className="step-summary">
        <div className="step-header" onClick={() => toggleStep('insights')}>
          <div className="step-info">
            <TrendingUp size={20} />
            <div>
              <h4>Cultural Popularity Insights</h4>
              <p>Qloo analyzed {job.popularity_insights.length} related cultural preferences</p>
            </div>
          </div>
          {isExpanded ? <ChevronUp size={20} /> : <ChevronDown size={20} />}
        </div>
        
        {isExpanded && (
          <div className="step-content">
            <div className="qloo-highlight">
              <Zap size={16} />
              <span>Qloo's Taste AI™ revealed what people actually enjoy and how interests connect</span>
            </div>
            <div className="insights-summary">
              <h5>Top Cultural Preferences:</h5>
              <div className="insights-list">
                {topInsights.map((insight, index) => (
                  <li key={index}>
                    <strong>{insight.name}</strong> ({insight.subtype}) - 
                    Popularity: {insight.popularity.toFixed(2)}
                  </li>
                ))}
              </div>
              {job.popularity_insights.length > 3 && (
                <p className="more-insights">+{job.popularity_insights.length - 3} more insights available</p>
              )}
            </div>
          </div>
        )}
      </div>
    );
  };

  const renderDemographicsStep = () => {
    if (!job.demographic_buckets || job.demographic_buckets.length === 0) return null;
    
    const isExpanded = expandedSteps.has('demographics');
    const totalBuckets = job.demographic_buckets.length;
    
    return (
      <div className="step-summary">
        <div className="step-header" onClick={() => toggleStep('demographics')}>
          <div className="step-info">
            <Users size={20} />
            <div>
              <h4>Audience Demographics</h4>
              <p>Qloo analyzed demographics for {totalBuckets} cultural entity{totalBuckets > 1 ? 'ies' : 'y'}</p>
            </div>
          </div>
          {isExpanded ? <ChevronUp size={20} /> : <ChevronDown size={20} />}
        </div>
        
        {isExpanded && (
          <div className="step-content">
            <div className="qloo-highlight">
              <Brain size={16} />
              <span>Qloo's privacy-first approach reveals audience insights without personal data</span>
            </div>
            {job.demographic_buckets.map((demo, index) => (
              <div key={index} className="demo-section">
                <h5>Entity: {demo.entity_id}</h5>
                
                {demo.age && Object.keys(demo.age).length > 0 && (
                  <div>
                    <h6>Age Distribution</h6>
                    <div className="demographics-grid">
                      {Object.entries(demo.age).map(([age, percentage]) => (
                        <div key={age} className="demo-item">
                          <h5>{age}</h5>
                          <p>{(percentage * 100).toFixed(1)}%</p>
                        </div>
                      ))}
                    </div>
                  </div>
                )}
                
                {demo.gender && Object.keys(demo.gender).length > 0 && (
                  <div>
                    <h6>Gender Distribution</h6>
                    <div className="demographics-grid">
                      {Object.entries(demo.gender).map(([gender, percentage]) => (
                        <div key={gender} className="demo-item">
                          <h5>{gender}</h5>
                          <p>{(percentage * 100).toFixed(1)}%</p>
                        </div>
                      ))}
                    </div>
                  </div>
                )}
              </div>
            ))}
          </div>
        )}
      </div>
    );
  };

  const renderJobInputs = () => {
    const isExpanded = expandedSteps.has('inputs');
    
    return (
      <div className="step-summary">
        <div className="step-header" onClick={() => toggleStep('inputs')}>
          <div className="step-info">
            <FileText size={20} />
            <div>
              <h4>Campaign Parameters</h4>
              <p>Product: {job.job_inputs.product} | Platform: {job.job_inputs.target_platform}</p>
            </div>
          </div>
          {isExpanded ? <ChevronUp size={20} /> : <ChevronDown size={20} />}
        </div>
        
        {isExpanded && (
          <div className="step-content">
            <div className="data-grid">
              <div className="data-item">
                <h5>Product</h5>
                <p>{job.job_inputs.product}</p>
              </div>
              <div className="data-item">
                <h5>Platform</h5>
                <p>{job.job_inputs.target_platform}</p>
              </div>
              <div className="data-item">
                <h5>Title</h5>
                <p>{job.job_inputs.title}</p>
              </div>
            </div>
          </div>
        )}
      </div>
    );
  };

  return (
    <div className="job-steps-section">
      <div className="steps-header">
        <h2>
          <Target size={24} />
           Internal insights process (Qloo + LLM)
        </h2>
        <div className="qloo-description">
          <Brain size={16} />
          <span>Qloo's Taste AI™ combines cultural knowledge with consumer behavior to understand how people interact with the world around them.</span>
        </div>
      </div>
      
      {renderJobInputs()}
      {renderSeedStep()}
      {renderSearchStep()}
      {renderInsightsStep()}
      {renderDemographicsStep()}
    </div>
  );
};

export default JobSteps; 