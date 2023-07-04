import React, { useState } from 'react';

function Image(
  props: React.DetailedHTMLProps<
    React.ImgHTMLAttributes<HTMLImageElement>,
    HTMLImageElement
  >
) {
  const [loaded, setLoaded] = useState(false);

  return (
    <div className={props.className}>
      {!loaded && <div style={{ width: '100%', height: '100%' }}></div>}
      <img
        {...props}
        onLoad={() => {
          setLoaded(true);
        }}
        style={loaded ? props.style : { display: 'none' }}
      />
    </div>
  );
}

export default Image;
