import axios from 'axios'
import React from 'react'
import { FaCloudDownloadAlt } from 'react-icons/fa'

export const SearchResult = ({ metadata, stateHandler }) => {

const importShow = () => {
    if (metadata !== null) {
       axios.post('http://10.0.0.220:10090/import/' + metadata.id)
       .then(function() {
        document.getElementById('importBtn').style = "display: none"
        stateHandler()
       })
       .catch(function(error) {
        console.log(error)
       })
    } else {
      console.log("No metadata found")
    }
}

return (
    <div className={`task ${metadata.status === 'Running' ? 
     'reminder' : ''}`}>
      <div className="row">
        <div className="column">
          <img alt={metadata.name} src={metadata.image.medium} width="150" height="200"></img>
        </div>
        <div className="column">
          <h3>{metadata.name}</h3>
          <button id="importBtn" className="nextBtn" onClick={() => importShow()}><FaCloudDownloadAlt/></button>
        </div>
      </div>
    </div>
  )
}
