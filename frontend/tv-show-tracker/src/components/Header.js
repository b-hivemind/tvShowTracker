import Button from './Button'

const Header = ({ title }) => {
    const onClick = () => {
        console.log('click')
    }

    return (
        <header className='header'>
            <h1>{title}</h1>
            <Button color='green' text='Import' onClick={onClick}/>
        </header>
    )
}

Header.defaultProps = {
    title: 'Tv Show Tracker',
}
export default Header