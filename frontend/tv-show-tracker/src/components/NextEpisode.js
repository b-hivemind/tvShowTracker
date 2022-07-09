const NextEpisode = ({ episode }) => {
    if (!episode) return null;
    return (
        <div className="nextEpisode">
            <h4>
                {episode.name} <br />
                Season {episode.season} <br />
                Episode {episode.number} <br />
                Serial {episode.serial} <br />
            </h4>
        </div>
    ) 
}

export default NextEpisode