import React from 'react';
import { Sparkles, Brain, Zap } from 'lucide-react';

const Header = () => {
  return (
    <div className="header">
      <h1>
        <Sparkles className="header-icon" />
        Qlaire
      </h1>
      <p>AI-Powered Ad Campaign Generator</p>
      <div className="header-subtitle">
        <div className="hackathon-badge">
          <Zap size={14} />
          <span>Qloo LLM Hackathon</span>
        </div>
      </div>
      <div className="header-description">
        <p>Combining cultural intelligence with LLMs to create campaigns that truly resonate with your audience.</p>
      </div>
    </div>
  );
};

export default Header; 