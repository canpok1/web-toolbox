import { Link, Outlet } from "react-router-dom";

const Layout = () => {
  return (
    <div>
      <header className="bg-neutral shadow">
        <div className="navbar container mx-auto text-neutral-content">
          <div className="flex-1">
            <Link to="/" className="btn btn-ghost text-xl normal-case">
              Web Toolbox
            </Link>
          </div>
          <div className="flex-none">
            <ul className="menu menu-horizontal px-1">
              <li>
                <a href="/planning-poker">プランニングポーカー</a>
              </li>
            </ul>
          </div>
        </div>
      </header>

      <main className="container mx-auto p-4">
        <Outlet />
      </main>
    </div>
  );
};

export default Layout;
