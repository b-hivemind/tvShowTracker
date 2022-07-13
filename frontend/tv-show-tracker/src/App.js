import { useState, useEffect } from 'react'
import axios from 'axios'

import Header from './components/Header'
import Shows from './components/Shows';
import Importer from './components/Importer';

function App() {
  const [shows, setShows] = useState(null)
  const [showImport, setShowImport] = useState(false)

  let getData = () => {
    axios.get("http://10.0.0.220:10090/shows").then((response) => {
      setShows(response.data)
    })
  }

  useEffect(() => {
    getData()
  }, [])


  return (
    <div className="container">
      <Header onImport={() => setShowImport(!showImport)} title="Tv Show Tracker" />
      {showImport && <Importer stateHandler={getData}/>}
      {shows != null && <Shows shows={shows} />}
    </div>
  );
}

export default App;
