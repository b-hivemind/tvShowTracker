import { useState } from 'react'
import React from 'react'
import axios from 'axios'
import { SearchResult } from './SearchResult'

const Importer = ({ stateHandler }) => {
  const [searchResults, setsearchResults] = useState(null)

  const searchShows = () => {  
    let query = document.getElementById('searchBox').value
    if (query.length > 0) {
      axios.post('http://10.0.0.220:10090/search', {
        'message': query
      })
      .then(function(response) {
        setsearchResults(response.data)
      })
      .catch(function(error) {
        console.log(error);
      })
    }
  } 
  
  return (
    <div className='form-control'>
        <div className="search-form">
          <label>Show</label>
          <input id='searchBox' type='text' placeholder='Search TVMaze'/>
          <button onClick={searchShows}>Find</button>
        </div>
        <div className='search-results'>
          {searchResults !== null && searchResults.map((result) => <SearchResult key={result.id} metadata={result} stateHandler={stateHandler}/>)}
        </div>
    </div>
  )
}

export default Importer;