const NewThemeLink = () => {
  return (
    <>
      <a
        href="https://forms.gle/your-google-form-url"
        target="_blank"
        rel="noopener noreferrer"
        className="mt-2 block font-semibold text-blue-600 text-sm transition-colors duration-200 hover:text-blue-800"
      >
        新しいトークテーマを投稿する
      </a>
      <p className="mt-1 text-gray-500 text-xs">
        良いテーマを思いついたら、ぜひ投稿してください！
      </p>
    </>
  );
};

export default NewThemeLink;
