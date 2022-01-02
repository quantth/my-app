import logo from './logo.svg';
import './styles/index.scss'
import { BrowserRouter as Router, Route, Routes} from 'react-router-dom'; 
import Navbar from './app/components/Navbar';
import About from './app/components/about/About';
import Posts from './app/components/posts/Posts';

function App() {
  return (
    <Router>
      <div className="App">
        <Navbar/>
        <div className="content">
          <Routes>
            <Route exact path='/' element={< Posts />}></Route>
            <Route exact path='/about' element={< About />}></Route>
          </Routes>
        </div>
      </div>
    </Router>
  );
}

export default App;
