const Button = ({ color, text, onClick }) => {
  return (
    <button style={{ backgroundColor: color }} 
     onClick={onClick}
     className='btn'>
      {text}
    </button>
    )
}

Button.defaultProps = {
  color: 'steelblue'
}

export default Button