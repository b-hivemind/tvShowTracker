import { useState, useEffect } from 'react'
import axios from 'axios'

import Header from './components/Header'
import Shows from './components/Shows';

function App() {
  const [shows, setShows] = useState(null)


  useEffect(() => {
    axios.get("http://10.0.0.220:10090/shows").then((response) => {
      setShows(response.data)
    })
  }, [])

  if (!shows) return null;

  return (
    <div className="container">
      <Header title="Tv Show Tracker"/>
      <Shows shows={shows}/>
    </div>
  );
}

export default App;
