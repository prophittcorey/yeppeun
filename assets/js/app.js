(function (d) {
  var cleaned = d.querySelector('code.language-js');
  var copy = d.querySelector('pre > .copy');

  if (cleaned && copy) {
    copy.addEventListener('click', function () {
      console.log('copied');
      navigator.clipboard.writeText(cleaned.innerText);
    });
  }

})(document)
