import React, { useState, useEffect } from 'react';
import axios from 'axios';
import Header from './components/Header';
import CampaignForm from './components/CampaignForm';
import JobStatus from './components/JobStatus';
import CampaignResults from './components/CampaignResults';
import JobSteps from './components/JobSteps';
import ServerStatus from './components/ServerStatus';
import './App.css';

const API_BASE_URL = process.env.REACT_APP_API_URL || '';

function App() {
  const [currentJob, setCurrentJob] = useState(null);
  const [pollingInterval, setPollingInterval] = useState(null);
  const [error, setError] = useState(null);
  const [showForm, setShowForm] = useState(true);

  const statusSteps = {
    'pending': 0,
    'generating_seed': 1,
    'searching_entity': 2,
    'fetching_insights': 3,
    'fetching_demographics': 4,
    'generating_output': 5,
    'completed': 6,
    'failed': -1
  };

  const startCampaign = async (campaignData) => {
    try {
      setError(null);
      const response = await axios.post(`${API_BASE_URL}/v1/api/ads/insights`, campaignData);
      const job = response.data.job;
      setCurrentJob(job);
      setShowForm(false);
      startPolling(job.id);
      return job;
    } catch (error) {
      const errorMessage = error.response?.data?.message || error.message || 'Failed to start campaign';
      setError(errorMessage);
      throw new Error(errorMessage);
    }
  };

  const startPolling = (jobId) => {
    if (pollingInterval) {
      clearInterval(pollingInterval);
    }

    const interval = setInterval(async () => {
      try {
        const response = await axios.get(`${API_BASE_URL}/v1/api/ads/insights/${jobId}`);
        const job = response.data.job;
        setCurrentJob(job);

        if (job.status === 'completed' || job.status === 'failed') {
          clearInterval(interval);
          setPollingInterval(null);
        }
      } catch (error) {
        console.error('Polling error:', error);
      }
    }, 2000);

    setPollingInterval(interval);
  };

  const getStatusMessage = (status) => {
    const messages = {
      'pending': 'Job created, starting processing...',
      'generating_seed': 'Generating Qloo search seed...',
      'searching_entity': 'Searching for relevant entities...',
      'fetching_insights': 'Fetching popularity insights...',
      'fetching_demographics': 'Analyzing demographics...',
      'generating_output': 'Generating final campaign...',
      'completed': 'Campaign generated successfully!',
      'failed': 'Job failed'
    };
    return messages[status] || status;
  };

  const resetForm = () => {
    setCurrentJob(null);
    setShowForm(true);
    setError(null);
    if (pollingInterval) {
      clearInterval(pollingInterval);
      setPollingInterval(null);
    }
  };

  useEffect(() => {
    return () => {
      if (pollingInterval) {
        clearInterval(pollingInterval);
      }
    };
  }, [pollingInterval]);

  return (
    <ServerStatus>
      <div className="container">
        <Header job={currentJob} />
        
        {error && (
          <div className="error-message">
            <i className="fas fa-exclamation-triangle"></i> {error}
          </div>
        )}

        {!showForm && currentJob && (
          <>
            <JobStatus 
              job={currentJob} 
              statusMessage={getStatusMessage(currentJob.status)}
              statusSteps={statusSteps}
              onReset={resetForm}
            />
            <JobSteps job={currentJob} />
          </>
        )}

        {showForm ? (
          <div className="form-container">
            <CampaignForm onSubmit={startCampaign} />
          </div>
        ) : (
          <div className="results-container">
            <CampaignResults job={currentJob} />
          </div>
        )}
      </div>
    </ServerStatus>
  );
}

export default App; 