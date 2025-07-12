import React from 'react';
import { Brain, Zap } from 'lucide-react';

const Header = ({ job }) => {
  const isCompleted = job?.status === 'completed';
  
  return (
    <div className="header">
      <h1>
        <div className={`logo-container ${isCompleted ? 'logo-container-completed' : ''}`}>
          <img src="/static/img/qlaire_icon.png" alt="Qlaire Icon" className={`header-icon ${isCompleted ? 'header-icon-completed' : ''}`} />
          <img src="/static/img/qlaire_logo_color.png" alt="Qlaire" className={`header-logo ${isCompleted ? 'header-logo-completed' : ''}`} />
        </div>
      </h1>
      <p>AI-Powered Ad Campaign Generator</p>
      <div className="header-subtitle">
        <div className="hackathon-badge">
          <span className="chip">
            <Zap size={14} style={{ verticalAlign: 'middle', marginRight: 4 }} />
            <a 
              href="https://qloo-hackathon.devpost.com/" 
              target="_blank" 
              rel="noopener noreferrer"
              className="hackathon-link"
            >
              Qloo LLM Hackathon
            </a>
          </span>
          <span className="chip" style={{ marginLeft: '0.5em' }}>
            <a
              href="https://devpost.com/marcoo-boschetti"
              target="_blank"
              rel="noopener noreferrer"
              className="hackathon-link"
            >
              Created by Marco
            </a>
          </span>
        </div>
      </div>
      <div className="header-description">
        <p>Combining cultural intelligence with LLMs to create campaigns that truly resonate with your audience.</p>
      </div>
    </div>
  );
};

export default Header; 