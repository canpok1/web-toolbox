import { useRef } from "react";
import { Link, Outlet } from "react-router-dom";

const Layout = () => {
  const detailsRef = useRef<HTMLDetailsElement>(null);

  const closeDropdown = () => {
    if (detailsRef.current) {
      detailsRef.current.open = false;
    }
  };

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
                <details ref={detailsRef} className="dropdown dropdown-end">
                  <summary className="btn btn-ghost">ツール一覧</summary>
                  <ul className="menu dropdown-content z-1 mt-4 w-52 rounded-box bg-base-100 p-2 shadow">
                    <li>
                      <Link
                        to="/planning-poker"
                        className="text-primary-content"
                        onClick={closeDropdown}
                      >
                        プランニングポーカー
                      </Link>
                    </li>
                  </ul>
                </details>
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
