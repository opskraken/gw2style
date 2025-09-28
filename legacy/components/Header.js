export default function Header() {
  return (
    <nav className="navbar">
      {/* Logo/Title */}
      <div className="navbar-logo">
        <a href="/">
          GW2🎨STYLE
        </a>
      </div>

      {/* Navigation Links */}
      <div className="navbar-links">
        <a href="/" className="navbar-link">
          Post
        </a>
        <a href="/search" className="navbar-link">
          Search
        </a>
        <a href="/about" className="navbar-link">
          About
        </a>
        <a href="/guidelines" className="navbar-link">
          Submission Guidelines
        </a>
        <a href="/login" className="navbar-link">
          Log In
        </a>
      </div>
    </nav>
  );
}