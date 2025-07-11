import React from 'react';
import { CheckCircle, Clock, AlertCircle, Loader, RefreshCw, Brain } from 'lucide-react';

const JobStatus = ({ job, statusMessage, statusSteps, onReset }) => {
  const getStatusIcon = (status) => {
    switch (status) {
      case 'completed':
        return <CheckCircle className="status-icon completed" />;
      case 'failed':
        return <AlertCircle className="status-icon failed" />;
      case 'pending':
        return <Clock className="status-icon pending" />;
      default:
        return <Loader className="status-icon processing" />;
    }
  };

  const getStatusColor = (status) => {
    switch (status) {
      case 'completed':
        return 'status-completed';
      case 'failed':
        return 'status-failed';
      case 'pending':
        return 'status-pending';
      default:
        return 'status-processing';
    }
  };

  const currentStep = statusSteps[job.status];
  const progress = (currentStep / 5) * 100;

  return (
    <div className="status-section">
      <div className="status-header">
        <h2>Job Status</h2>
        <div className="qloo-badge">
          <Brain size={16} />
          Powered by Qloo Taste AI™
        </div>
      </div>
      
      <div className="status-indicator">
        {getStatusIcon(job.status)}
        <span className={getStatusColor(job.status)}>{statusMessage}</span>
      </div>
      
      <div className="progress-bar">
        <div 
          className="progress-fill" 
          style={{ width: `${Math.min(progress, 100)}%` }}
        />
      </div>
      
      <div className="step-indicators">
        {[
          { id: 1, label: 'Generate Seed', step: 'generating_seed' },
          { id: 2, label: 'Search Entity', step: 'searching_entity' },
          { id: 3, label: 'Fetch Insights', step: 'fetching_insights' },
          { id: 4, label: 'Demographics', step: 'fetching_demographics' },
          { id: 5, label: 'Generate Campaign', step: 'generating_output' }
        ].map(({ id, label, step }) => {
          const stepNumber = statusSteps[step];
          let className = 'step';
          
          if (stepNumber < currentStep) {
            className += ' completed';
          } else if (stepNumber === currentStep) {
            className += ' active';
          }

          return (
            <div key={id} className={className}>
              <div className="step-icon">{id}</div>
              <div className="step-label">{label}</div>
            </div>
          );
        })}
      </div>

      {job.final_error && (
        <div className="error-message">
          <AlertCircle />
          {job.final_error}
        </div>
      )}

      {(job.status === 'completed' || job.status === 'failed') && (
        <div className="status-actions">
          <button onClick={onReset} className="btn-secondary">
            <RefreshCw size={16} />
            Create New Campaign
          </button>
        </div>
      )}
    </div>
  );
};

export default JobStatus; 