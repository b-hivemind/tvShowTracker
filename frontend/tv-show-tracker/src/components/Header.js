import Button from './Button'

const Header = ({ title, onImport }) => {
    return (
        <header className='header'>
            <h1>{title}</h1>
            <Button color='green' text='Import' onClick={onImport}/>
        </header>
    )
}

Header.defaultProps = {
    title: 'Tv Show Tracker',
}
export default Header