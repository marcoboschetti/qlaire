import React from 'react';
import { Sparkles } from 'lucide-react';

const Header = () => {
  return (
    <div className="header">
      <h1>
        <Sparkles className="header-icon" />
        Qlaire
      </h1>
      <p>AI-Powered Ad Campaign Generator</p>
    </div>
  );
};

export default Header; 