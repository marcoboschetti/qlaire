import React, { useState } from 'react';
import { Rocket, Loader } from 'lucide-react';

const CampaignForm = ({ onSubmit, disabled }) => {
  const [formData, setFormData] = useState({
    product: '',
    target_platform: '',
    title: ''
  });
  const [isSubmitting, setIsSubmitting] = useState(false);

  const handleSubmit = async (e) => {
    e.preventDefault();
    setIsSubmitting(true);
    
    try {
      await onSubmit(formData);
      setFormData({ product: '', target_platform: '', title: '' });
    } catch (error) {
      console.error('Form submission error:', error);
    } finally {
      setIsSubmitting(false);
    }
  };

  const handleChange = (e) => {
    setFormData({
      ...formData,
      [e.target.name]: e.target.value
    });
  };

  return (
    <div className="form-section">
      <h2>
        <i className="fas fa-edit"></i> Campaign Details
      </h2>
      <form onSubmit={handleSubmit}>
        <div className="form-group">
          <label htmlFor="product">Product Name *</label>
          <input
            type="text"
            id="product"
            name="product"
            value={formData.product}
            onChange={handleChange}
            required
            placeholder="e.g., Premium Wireless Headphones"
            disabled={disabled || isSubmitting}
          />
        </div>
        
        <div className="form-group">
          <label htmlFor="platform">Target Platform *</label>
          <select
            id="platform"
            name="target_platform"
            value={formData.target_platform}
            onChange={handleChange}
            required
            disabled={disabled || isSubmitting}
          >
            <option value="">Select a platform</option>
            <option value="Meta Ads">Meta Ads (Facebook/Instagram)</option>
            <option value="Google Ads">Google Ads</option>
            <option value="TikTok Ads">TikTok Ads</option>
            <option value="LinkedIn Ads">LinkedIn Ads</option>
            <option value="Twitter Ads">Twitter Ads</option>
          </select>
        </div>
        
        <div className="form-group">
          <label htmlFor="title">Campaign Title *</label>
          <input
            type="text"
            id="title"
            name="title"
            value={formData.title}
            onChange={handleChange}
            required
            placeholder="e.g., Summer Sale - Premium Audio Experience"
            disabled={disabled || isSubmitting}
          />
        </div>
        
        <button 
          type="submit" 
          className="btn" 
          disabled={disabled || isSubmitting}
        >
          {isSubmitting ? (
            <>
              <Loader className="spinner" />
              Starting...
            </>
          ) : (
            <>
              <Rocket />
              Generate Campaign
            </>
          )}
        </button>
      </form>
    </div>
  );
};

export default CampaignForm; 