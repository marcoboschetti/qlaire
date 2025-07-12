import React, { useState, useEffect } from 'react';
import { Loader2, AlertCircle, X } from 'lucide-react';

const ServerStatus = ({ children }) => {
  const [serverStatus, setServerStatus] = useState('checking'); // 'checking', 'online', 'offline'
  const [retryCount, setRetryCount] = useState(0);
  const [showToast, setShowToast] = useState(true);
  const maxRetries = 10;

  const checkServerHealth = async () => {
    try {
      const response = await fetch('/ping', {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        },
      });
      
      if (response.ok) {
        setServerStatus('online');
        // Hide toast after a short delay when server is online
        setTimeout(() => setShowToast(false), 2000);
      } else {
        throw new Error('Server responded with error');
      }
    } catch (error) {
      console.log('Server health check failed:', error);
      if (retryCount < maxRetries) {
        setRetryCount(prev => prev + 1);
        // Retry after 2 seconds
        setTimeout(checkServerHealth, 2000);
      } else {
        setServerStatus('offline');
      }
    }
  };

  useEffect(() => {
    checkServerHealth();
  }, []);

  useEffect(() => {
    if (serverStatus === 'checking' && retryCount < maxRetries) {
      const timer = setTimeout(checkServerHealth, 2000);
      return () => clearTimeout(timer);
    }
  }, [retryCount, serverStatus]);

  return (
    <>
      {children}
      
      {/* Non-blocking toast notification */}
      {showToast && serverStatus === 'checking' && (
        <div className="server-toast">
          <div className="server-toast-content">
            <div className="server-toast-header">
              <Loader2 className="server-toast-spinner" size={16} />
              <span>Starting Qlaire Server...</span>
              <button 
                className="server-toast-close"
                onClick={() => setShowToast(false)}
              >
                <X size={14} />
              </button>
            </div>
            <div className="server-toast-disclaimer">
              <AlertCircle size={12} />
              <span>
                <strong>Hackathon Project:</strong> Running on suboptimal infrastructure. 
                Thank you for your patience! (Attempt {retryCount + 1}/{maxRetries})
              </span>
            </div>
          </div>
        </div>
      )}

      {/* Error toast */}
      {showToast && serverStatus === 'offline' && (
        <div className="server-toast server-toast-error">
          <div className="server-toast-content">
            <div className="server-toast-header">
              <AlertCircle className="server-toast-error-icon" size={16} />
              <span>Server Unavailable</span>
              <button 
                className="server-toast-close"
                onClick={() => setShowToast(false)}
              >
                <X size={14} />
              </button>
            </div>
            <div className="server-toast-disclaimer">
              <span>
                The server may take longer to start. You can continue using the form, 
                but submissions will fail until the server is ready.
              </span>
            </div>
            <button 
              className="server-toast-retry-btn" 
              onClick={() => {
                setServerStatus('checking');
                setRetryCount(0);
                checkServerHealth();
              }}
            >
              Retry Connection
            </button>
          </div>
        </div>
      )}
    </>
  );
};

export default ServerStatus; 