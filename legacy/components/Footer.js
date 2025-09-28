export default function Footer() {
  return (
    <footer className="footer">
      <div className="footer-content">
        {/* Main sections */}
        <div className="footer-section">
          <h3 className="footer-title">GW2🎨STYLE</h3>
          <p className="footer-description">
            Share your Guild Wars 2 fashion creations with the community. 
            Made by players, for players.
          </p>
        </div>

        <div className="footer-section">
          <h4 className="footer-heading">Browse</h4>
          <ul className="footer-links">
            <li><a href="/gallery">Gallery</a></li>
            <li><a href="/popular">Popular Looks</a></li>
            <li><a href="/recent">Recent Posts</a></li>
            <li><a href="/dyes">Dye Guides</a></li>
          </ul>
        </div>

        <div className="footer-section">
          <h4 className="footer-heading">Community</h4>
          <ul className="footer-links">
            <li><a href="/post">Submit Look</a></li>
            <li><a href="/guidelines">Guidelines</a></li>
            <li><a href="/contests">Contests</a></li>
            <li><a href="/about">About Us</a></li>
          </ul>
        </div>

        <div className="footer-section">
          <h4 className="footer-heading">Guild Wars 2</h4>
          <ul className="footer-links">
            <li><a href="https://guildwars2.com" target="_blank" rel="noopener noreferrer">Official Site</a></li>
            <li><a href="https://wiki.guildwars2.com" target="_blank" rel="noopener noreferrer">GW2 Wiki</a></li>
            <li><a href="#" className="social-link">Reddit 🔗</a></li>
            <li><a href="#" className="social-link">Discord 🔗</a></li>
          </ul>
        </div>
      </div>

      {/* Bottom section */}
      <div className="footer-bottom">
        <div className="footer-bottom-content">
          <p className="footer-copyright">
            GW2🎨STYLE is an independent fan project.
          </p>
          <div className="footer-bottom-links">
            <a href="/privacy">Privacy</a>
            <a href="/terms">Terms</a>
            <a href="/contact">Contact</a>
          </div>
        </div>
      </div>
    </footer>
  );
}
