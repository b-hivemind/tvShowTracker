import axios from 'axios'
import { useState, useEffect } from 'react'
import { FaAngleDoubleRight } from 'react-icons/fa'

import NextEpisode from './NextEpisode'

const Show = ({ show }) => {
  const [nextEp, setNextEp] = useState(null)

  useEffect(() => {
    const getNextEp = async (show) => {
      const queryURL = "http://10.0.0.220:10090/shows/" + show.id + "/next"
      const results = await axios.get(queryURL)
      setNextEp(results.data)
    }
    getNextEp(show)
  }, [])

  const setNext = () => {
    const setNextEpisode = async (show) => {
      const queryURL = "http://10.0.0.220:10090/shows/" + show.id + "/next"
      const results = await axios.post(queryURL)
      console.log(results.data)
      setNextEp(results.data)
    }
    setNextEpisode(show)
  }

  return (
   <div className={`task ${show.status === 'Running' ? 
   'reminder' : ''}`}>
      <div className="row">
        <div className="column">
          <img src={show.image.medium} width="150" height="200"></img>
        </div>
        <div className="column">
          <h3>{show.name}</h3>
          <NextEpisode episode={nextEp}/>
          <button className="nextBtn" onClick={() => setNext()}><FaAngleDoubleRight/></button>
        </div>
      </div>
    </div>
  )
}

export default Show

