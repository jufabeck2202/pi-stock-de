interface LogoProps extends React.ComponentPropsWithoutRef<"svg"> {
  width?: number;
}
export function Logo({ width = 75, ...others }: LogoProps) {
  return (
    <svg
      viewBox="0 0 314 196"
      {...others}
      xmlns="http://www.w3.org/2000/svg"
      width={width}
    >
      <rect width="302" height="185" rx="30" fill="#50515D" />
      <circle cx="28" cy="28" r="15" fill="#C4C4C4" />
      <circle cx="28" cy="28" r="6" fill="#6C6C6C" />
      <circle cx="182" cy="28" r="6" fill="#6C6C6C" />
      <circle cx="67" cy="28" r="6" fill="#6C6C6C" />
      <circle cx="28" cy="158" r="15" fill="#C4C4C4" />
      <circle cx="28" cy="158" r="6" fill="#6C6C6C" />
      <circle cx="216" cy="158" r="15" fill="#C4C4C4" />
      <circle cx="216" cy="158" r="6" fill="#6C6C6C" />
      <circle cx="216" cy="28" r="15" fill="#C4C4C4" />
      <circle cx="216" cy="28" r="6" fill="#6C6C6C" />
      <g filter="url(#filter0_d_4_4)">
        <rect x="257" y="15" width="53" height="39" rx="4" fill="#C4C4C4" />
      </g>
      <g filter="url(#filter1_d_4_4)">
        <rect x="246" y="111" width="64" height="47" rx="5" fill="#C4C4C4" />
      </g>
      <g filter="url(#filter2_d_4_4)">
        <rect x="257" y="63" width="53" height="39" rx="4" fill="#C4C4C4" />
      </g>
      <g filter="url(#filter3_d_4_4)">
        <rect x="91" y="164" width="47" height="24" rx="4" fill="#C4C4C4" />
      </g>
      <g filter="url(#filter4_d_4_4)">
        <rect x="55" y="164" width="24" height="24" rx="4" fill="#C4C4C4" />
      </g>
      <g filter="url(#filter5_d_4_4)">
        <rect x="91" y="63" width="60" height="60" rx="4" fill="#C4C4C4" />
      </g>
      <rect x="67" y="22" width="115" height="12" fill="#6C6C6C" />
      <circle
        cx="29.5"
        cy="72.5"
        r="7.5"
        transform="rotate(-90 29.5 72.5)"
        fill="#6C6C6C"
      />
      <circle
        cx="29.5"
        cy="113.5"
        r="7.5"
        transform="rotate(-90 29.5 113.5)"
        fill="#6C6C6C"
      />
      <rect
        x="22"
        y="113"
        width="40"
        height="15"
        transform="rotate(-90 22 113)"
        fill="#6C6C6C"
      />
      <defs>
        <filter
          id="filter0_d_4_4"
          x="253"
          y="15"
          width="61"
          height="47"
          filterUnits="userSpaceOnUse"
          color-interpolation-filters="sRGB"
        >
          <feFlood flood-opacity="0" result="BackgroundImageFix" />
          <feColorMatrix
            in="SourceAlpha"
            type="matrix"
            values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0"
            result="hardAlpha"
          />
          <feOffset dy="4" />
          <feGaussianBlur stdDeviation="2" />
          <feComposite in2="hardAlpha" operator="out" />
          <feColorMatrix
            type="matrix"
            values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0.25 0"
          />
          <feBlend
            mode="normal"
            in2="BackgroundImageFix"
            result="effect1_dropShadow_4_4"
          />
          <feBlend
            mode="normal"
            in="SourceGraphic"
            in2="effect1_dropShadow_4_4"
            result="shape"
          />
        </filter>
        <filter
          id="filter1_d_4_4"
          x="242"
          y="111"
          width="72"
          height="55"
          filterUnits="userSpaceOnUse"
          color-interpolation-filters="sRGB"
        >
          <feFlood flood-opacity="0" result="BackgroundImageFix" />
          <feColorMatrix
            in="SourceAlpha"
            type="matrix"
            values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0"
            result="hardAlpha"
          />
          <feOffset dy="4" />
          <feGaussianBlur stdDeviation="2" />
          <feComposite in2="hardAlpha" operator="out" />
          <feColorMatrix
            type="matrix"
            values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0.25 0"
          />
          <feBlend
            mode="normal"
            in2="BackgroundImageFix"
            result="effect1_dropShadow_4_4"
          />
          <feBlend
            mode="normal"
            in="SourceGraphic"
            in2="effect1_dropShadow_4_4"
            result="shape"
          />
        </filter>
        <filter
          id="filter2_d_4_4"
          x="253"
          y="63"
          width="61"
          height="47"
          filterUnits="userSpaceOnUse"
          color-interpolation-filters="sRGB"
        >
          <feFlood flood-opacity="0" result="BackgroundImageFix" />
          <feColorMatrix
            in="SourceAlpha"
            type="matrix"
            values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0"
            result="hardAlpha"
          />
          <feOffset dy="4" />
          <feGaussianBlur stdDeviation="2" />
          <feComposite in2="hardAlpha" operator="out" />
          <feColorMatrix
            type="matrix"
            values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0.25 0"
          />
          <feBlend
            mode="normal"
            in2="BackgroundImageFix"
            result="effect1_dropShadow_4_4"
          />
          <feBlend
            mode="normal"
            in="SourceGraphic"
            in2="effect1_dropShadow_4_4"
            result="shape"
          />
        </filter>
        <filter
          id="filter3_d_4_4"
          x="87"
          y="164"
          width="55"
          height="32"
          filterUnits="userSpaceOnUse"
          color-interpolation-filters="sRGB"
        >
          <feFlood flood-opacity="0" result="BackgroundImageFix" />
          <feColorMatrix
            in="SourceAlpha"
            type="matrix"
            values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0"
            result="hardAlpha"
          />
          <feOffset dy="4" />
          <feGaussianBlur stdDeviation="2" />
          <feComposite in2="hardAlpha" operator="out" />
          <feColorMatrix
            type="matrix"
            values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0.25 0"
          />
          <feBlend
            mode="normal"
            in2="BackgroundImageFix"
            result="effect1_dropShadow_4_4"
          />
          <feBlend
            mode="normal"
            in="SourceGraphic"
            in2="effect1_dropShadow_4_4"
            result="shape"
          />
        </filter>
        <filter
          id="filter4_d_4_4"
          x="51"
          y="164"
          width="32"
          height="32"
          filterUnits="userSpaceOnUse"
          color-interpolation-filters="sRGB"
        >
          <feFlood flood-opacity="0" result="BackgroundImageFix" />
          <feColorMatrix
            in="SourceAlpha"
            type="matrix"
            values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0"
            result="hardAlpha"
          />
          <feOffset dy="4" />
          <feGaussianBlur stdDeviation="2" />
          <feComposite in2="hardAlpha" operator="out" />
          <feColorMatrix
            type="matrix"
            values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0.25 0"
          />
          <feBlend
            mode="normal"
            in2="BackgroundImageFix"
            result="effect1_dropShadow_4_4"
          />
          <feBlend
            mode="normal"
            in="SourceGraphic"
            in2="effect1_dropShadow_4_4"
            result="shape"
          />
        </filter>
        <filter
          id="filter5_d_4_4"
          x="87"
          y="63"
          width="68"
          height="68"
          filterUnits="userSpaceOnUse"
          color-interpolation-filters="sRGB"
        >
          <feFlood flood-opacity="0" result="BackgroundImageFix" />
          <feColorMatrix
            in="SourceAlpha"
            type="matrix"
            values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0"
            result="hardAlpha"
          />
          <feOffset dy="4" />
          <feGaussianBlur stdDeviation="2" />
          <feComposite in2="hardAlpha" operator="out" />
          <feColorMatrix
            type="matrix"
            values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0.25 0"
          />
          <feBlend
            mode="normal"
            in2="BackgroundImageFix"
            result="effect1_dropShadow_4_4"
          />
          <feBlend
            mode="normal"
            in="SourceGraphic"
            in2="effect1_dropShadow_4_4"
            result="shape"
          />
        </filter>
      </defs>
    </svg>
  );
}
