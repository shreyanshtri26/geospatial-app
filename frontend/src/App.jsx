import React, { useState } from 'react';
import Map from './component/Map';
import FileUpload from './component/FileUpload';
import './App.css';

const App = () => {
  const [geojsonData, setGeojsonData] = useState(null);
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');

  const handleFileUpload = (fileContent) => {
    setGeojsonData(fileContent);
  };

  const handleRegister = () => {
    // Send registration request to backend
    console.log('Register:', username, password);
  };

  const handleLogin = () => {
    // Send login request to backend
    console.log('Login:', username, password);
  };

  return (
    <div className="App">
      <div style={{ textAlign: 'center' }}>
        <h1>Geospatial App</h1>
        <div className="auth-container">
          <input
            type="text"
            placeholder="Username"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
          />
          <input
            type="password"
            placeholder="Password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
          />
          <button className="btn" onClick={handleRegister}>
            Create Account
          </button>
          <button className="btn" onClick={handleLogin}>
            Login
          </button>
        </div>
      </div>
      <div className="map-container">
        <FileUpload onFileUpload={handleFileUpload} />
        <Map geojsonData={geojsonData} />
      </div>
    </div>
  );
};

export default App;
