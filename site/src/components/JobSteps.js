import React from 'react';
import { Search, TrendingUp, Users, FileText, Target } from 'lucide-react';

const JobSteps = ({ job }) => {
  const renderSeedStep = () => {
    if (!job.generated_seed) return null;
    
    return (
      <div className="step-content">
        <h4>
          <Search size={20} />
          Generated Seed
        </h4>
        <div className="data-grid">
          <div className="data-item">
            <h5>Query</h5>
            <p>{job.generated_seed.query}</p>
          </div>
          <div className="data-item">
            <h5>Type</h5>
            <p>{job.generated_seed.type}</p>
          </div>
        </div>
      </div>
    );
  };

  const renderSearchStep = () => {
    if (!job.search_results || job.search_results.length === 0) return null;
    
    return (
      <div className="step-content">
        <h4>
          <Search size={20} />
          Search Results ({job.search_results.length} entities found)
        </h4>
        <div className="data-grid">
          {job.search_results.map((result, index) => (
            <div key={index} className="data-item">
              <h5>{result.name}</h5>
              <p><strong>ID:</strong> {result.entity_id}</p>
              <p><strong>Types:</strong> {result.types.join(', ')}</p>
              {result.short_desc && (
                <p><strong>Description:</strong> {result.short_desc}</p>
              )}
            </div>
          ))}
        </div>
      </div>
    );
  };

  const renderInsightsStep = () => {
    if (!job.popularity_insights || job.popularity_insights.length === 0) return null;
    
    return (
      <div className="step-content">
        <h4>
          <TrendingUp size={20} />
          Popularity Insights ({job.popularity_insights.length} insights)
        </h4>
        <div className="insights-list">
          {job.popularity_insights
            .sort((a, b) => b.popularity - a.popularity)
            .map((insight, index) => (
              <li key={index}>
                <strong>{insight.name}</strong> ({insight.subtype}) - 
                Popularity: {insight.popularity.toFixed(2)}
              </li>
            ))}
        </div>
      </div>
    );
  };

  const renderDemographicsStep = () => {
    if (!job.demographic_buckets || job.demographic_buckets.length === 0) return null;
    
    return (
      <div className="step-content">
        <h4>
          <Users size={20} />
          Demographic Analysis
        </h4>
        {job.demographic_buckets.map((demo, index) => (
          <div key={index} className="step-content">
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
    );
  };

  const renderJobInputs = () => {
    return (
      <div className="step-content">
        <h4>
          <FileText size={20} />
          Job Inputs
        </h4>
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
    );
  };

  return (
    <div className="job-steps-section">
      <h2>
        <Target size={24} />
        Processing Details
      </h2>
      
      {renderJobInputs()}
      {renderSeedStep()}
      {renderSearchStep()}
      {renderInsightsStep()}
      {renderDemographicsStep()}
    </div>
  );
};

export default JobSteps; 