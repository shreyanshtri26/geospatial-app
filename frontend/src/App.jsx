import React, { useState } from 'react';
import Map from './component/Map';
import FileUpload from './component/FileUpload';
import './App.css';

const App = () => {
  const [geojsonData, setGeojsonData] = useState(null);

  const handleFileUpload = (fileContent) => {
    setGeojsonData(fileContent);
  };

  return (
    <div className="App">
      <div style={{ textAlign: 'center' }}>
        <h1>Geospatial App</h1>
      </div>
      <div className="map-container">
        <FileUpload onFileUpload={handleFileUpload} />
        <Map geojsonData={geojsonData} />
      </div>
    </div>
  );
};

export default App;
