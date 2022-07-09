import Show from './Show'

const Shows = ({ shows }) => {
    return (
        <div className="showContainer">
            {shows.map((show) => <Show key={show.id} show={show}/>)}    
        </div>
    )
}

export default Shows
