import { Link } from 'react-router-dom';

const Navbar = () => {
    return (
        <nav className="navbar">
            <h1 class="smallcap">
                <a class="site--title" href="/">QUAN NGUYEN</a>
            </h1>
            <div class="links">
                <Link to="/">cd~</Link>
                <span> / </span>
                <Link to="/about">about</Link>
                <span> / </span>
                <Link to="/tags">tags</Link>
                <span> / </span>
                <Link to="/game/tic-tac-toe">game</Link>
                <span> / </span>
                <Link to="/rss">rss</Link>
            </div>
            <div class="intro">
                <p class="site-intro">Sometimes everything is so beautiful</p>
            </div>
            <div class="in--touch">
                <span>Find me on</span>
                <i class="fa fa-github fa-1x" aria-hidden="true"></i>
                <a href="https://github.com/qunv">GitHub</a>
                <i class="fa fa-twitter fa-1x" aria-hidden="true"></i>
                <a href="https://twitter.com/quannv132">Twitter</a>
            </div>
        </nav>
    );
}

export default Navbar;