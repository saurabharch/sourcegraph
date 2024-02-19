import type { FC } from 'react'

import styles from './SearchUpsellPage.module.scss'

interface BatchChangesLogoProps {
    className: string
}

export const BatchChangesLogo: FC<BatchChangesLogoProps> = ({ className }) => (
    <svg
        xmlns="http://www.w3.org/2000/svg"
        width="182"
        height="26"
        fill="none"
        viewBox="0 0 182 26"
        className={className}
    >
        <g clipPath="url(#clip0_990_5232)">
            <path
                stroke="url(#paint0_linear_990_5232)"
                strokeLinecap="round"
                strokeLinejoin="round"
                strokeWidth="2"
                d="M2.617 12h7.85m-7.85-8.572h7.85m0 17.143H4.361c-.963 0-1.744-.767-1.744-1.714v-18M20.062 14.57h-7.85c-.964 0-1.745-.767-1.745-1.714v-1.714c0-.947.781-1.715 1.745-1.715h7.85c.963 0 1.744.768 1.744 1.715v1.714c0 .947-.78 1.714-1.744 1.714zm0-8.571h-7.85c-.964 0-1.745-.768-1.745-1.715V2.571c0-.947.781-1.714 1.745-1.714h7.85c.963 0 1.744.767 1.744 1.714v1.714c0 .947-.78 1.715-1.744 1.715zm0 17.143h-7.85c-.964 0-1.745-.768-1.745-1.715v-1.714c0-.947.781-1.714 1.745-1.714h7.85c.963 0 1.744.767 1.744 1.714v1.714c0 .947-.78 1.715-1.744 1.715z"
            />
        </g>
        <path
            className={styles.otherIntegrationsLogoText}
            d="M178.515 17.188c0-.32-.139-.576-.416-.768a5.006 5.006 0 00-1.025-.528c-.407-.16-.855-.32-1.344-.48a5.604 5.604 0 01-1.343-.696 4.05 4.05 0 01-1.026-1.128c-.276-.464-.415-1.056-.415-1.776 0-1.184.358-2.088 1.075-2.712.716-.624 1.75-.936 3.101-.936.929 0 1.767.096 2.516.288.749.192 1.335.408 1.759.648l-.709 2.256a16.303 16.303 0 00-1.416-.456 6.387 6.387 0 00-1.734-.24c-.945 0-1.417.36-1.417 1.08 0 .288.139.52.415.696.277.176.619.344 1.026.504.407.144.855.304 1.343.48.489.176.937.408 1.344.696.407.272.749.632 1.025 1.08.277.448.416 1.024.416 1.728 0 1.216-.399 2.176-1.197 2.88-.782.688-1.962 1.032-3.541 1.032a8.657 8.657 0 01-2.443-.336c-.749-.208-1.359-.456-1.832-.744l.88-2.328c.374.208.871.416 1.49.624a6.384 6.384 0 001.905.288c.472 0 .846-.088 1.123-.264.293-.176.44-.472.44-.888zM171.335 19.468c-.489.384-1.156.712-2.003.984a8.989 8.989 0 01-2.662.384c-1.97 0-3.411-.56-4.323-1.68-.912-1.136-1.368-2.688-1.368-4.656 0-2.112.513-3.696 1.539-4.752 1.026-1.056 2.467-1.584 4.323-1.584.619 0 1.221.08 1.807.24.586.16 1.107.424 1.563.792.456.368.823.864 1.099 1.488.277.624.416 1.4.416 2.328 0 .336-.025.696-.074 1.08-.032.384-.089.784-.171 1.2h-7.327c.049 1.008.31 1.768.782 2.28.488.512 1.27.768 2.345.768.667 0 1.261-.096 1.782-.288.538-.208.945-.416 1.222-.624l1.05 2.04zm-4.543-8.808c-.83 0-1.449.248-1.856.744-.391.48-.627 1.128-.708 1.944h4.542c.065-.864-.073-1.528-.415-1.992-.326-.464-.847-.696-1.563-.696zM159.311 20.5c0 1.744-.48 3.024-1.441 3.84-.96.832-2.304 1.248-4.03 1.248-1.172 0-2.1-.08-2.784-.24-.667-.16-1.172-.328-1.514-.504l.659-2.472c.375.144.806.288 1.295.432.505.144 1.123.216 1.856.216 1.107 0 1.856-.24 2.247-.72.407-.464.61-1.12.61-1.968v-.768h-.097c-.57.752-1.58 1.128-3.029 1.128-1.579 0-2.76-.48-3.541-1.44-.765-.96-1.148-2.464-1.148-4.512 0-2.144.521-3.768 1.563-4.872 1.042-1.104 2.556-1.656 4.543-1.656 1.042 0 1.97.072 2.784.216.83.144 1.506.312 2.027.504V20.5zm-5.373-2.304c.619 0 1.091-.136 1.417-.408.342-.272.602-.68.781-1.224v-5.4c-.505-.208-1.131-.312-1.88-.312-.815 0-1.45.304-1.905.912-.456.592-.684 1.544-.684 2.856 0 1.168.195 2.056.586 2.664.391.608.952.912 1.685.912zM143.654 20.5v-6.816c0-.976-.147-1.68-.44-2.112-.277-.432-.757-.648-1.441-.648-.602 0-1.115.176-1.538.528-.407.336-.7.76-.88 1.272V20.5h-3.175v-12h2.516l.366 1.584h.098c.375-.512.871-.96 1.49-1.344.619-.384 1.416-.576 2.393-.576.603 0 1.14.08 1.612.24.472.16.871.424 1.197.792.326.368.57.872.733 1.512.162.624.244 1.4.244 2.328V20.5h-3.175zM124.5 9.196c.651-.288 1.424-.512 2.32-.672a14.537 14.537 0 012.809-.264c.846 0 1.555.104 2.124.312.57.192 1.018.472 1.344.84.342.368.578.808.708 1.32a6.26 6.26 0 01.22 1.728c0 .704-.025 1.416-.073 2.136-.049.704-.082 1.4-.098 2.088 0 .688.024 1.36.073 2.016.049.64.171 1.248.366 1.824h-2.588l-.513-1.656h-.122c-.326.496-.782.928-1.368 1.296-.57.352-1.311.528-2.223.528-.569 0-1.082-.08-1.538-.24a3.597 3.597 0 01-1.173-.72 3.452 3.452 0 01-.757-1.104 3.728 3.728 0 01-.268-1.44c0-.736.162-1.352.488-1.848.342-.512.822-.92 1.441-1.224.635-.32 1.384-.536 2.247-.648.879-.128 1.856-.168 2.931-.12.114-.896.049-1.536-.196-1.92-.244-.4-.789-.6-1.636-.6-.635 0-1.311.064-2.027.192-.7.128-1.278.296-1.734.504l-.757-2.328zm4.03 8.928c.635 0 1.139-.136 1.514-.408.374-.288.651-.592.83-.912v-1.56a9.02 9.02 0 00-1.465-.024 5.352 5.352 0 00-1.221.216c-.359.112-.643.272-.855.48a1.062 1.062 0 00-.318.792c0 .448.131.8.391 1.056.277.24.651.36 1.124.36zM119.247 20.5v-6.816c0-.976-.139-1.68-.416-2.112-.276-.432-.781-.648-1.514-.648-.537 0-1.034.184-1.49.552a2.528 2.528 0 00-.879 1.368V20.5h-3.175V3.7h3.175v6.144h.098c.391-.512.871-.92 1.441-1.224.57-.304 1.294-.456 2.174-.456.618 0 1.164.08 1.636.24.472.16.863.424 1.172.792.326.368.562.872.708 1.512.163.624.245 1.4.245 2.328V20.5h-3.175zM110.289 19.804c-.521.368-1.188.632-2.003.792-.797.16-1.62.24-2.466.24a9.355 9.355 0 01-2.931-.456 6.467 6.467 0 01-2.418-1.512c-.7-.704-1.262-1.608-1.685-2.712-.407-1.12-.61-2.472-.61-4.056 0-1.648.227-3.024.683-4.128.472-1.12 1.075-2.016 1.807-2.688a6.593 6.593 0 012.492-1.464 8.596 8.596 0 012.735-.456c.977 0 1.799.064 2.467.192.684.128 1.245.28 1.685.456l-.659 2.784a4.925 4.925 0 00-1.344-.384c-.505-.08-1.123-.12-1.856-.12-1.351 0-2.442.472-3.273 1.416-.814.944-1.221 2.408-1.221 4.392 0 .864.098 1.656.293 2.376.196.704.489 1.312.88 1.824.407.496.903.888 1.489 1.176.603.272 1.295.408 2.076.408.733 0 1.352-.072 1.856-.216a5.49 5.49 0 001.319-.552l.684 2.688zM87.955 20.5v-6.816c0-.976-.139-1.68-.416-2.112-.276-.432-.781-.648-1.514-.648-.537 0-1.034.184-1.49.552-.44.352-.732.808-.879 1.368V20.5h-3.175V3.7h3.175v6.144h.098c.39-.512.87-.92 1.44-1.224.57-.304 1.295-.456 2.174-.456.62 0 1.165.08 1.637.24.472.16.863.424 1.172.792.326.368.562.872.708 1.512.163.624.245 1.4.245 2.328V20.5h-3.175zM79.07 19.756c-.489.352-1.083.616-1.783.792a7.823 7.823 0 01-2.125.288c-.977 0-1.807-.152-2.491-.456a4.281 4.281 0 01-1.637-1.272c-.423-.56-.732-1.232-.928-2.016a11.647 11.647 0 01-.268-2.592c0-2.032.464-3.592 1.392-4.68.928-1.104 2.28-1.656 4.054-1.656.896 0 1.612.072 2.15.216.553.144 1.05.328 1.49.552l-.758 2.544a6.346 6.346 0 00-1.123-.408 4.835 4.835 0 00-1.246-.144c-.88 0-1.547.288-2.003.864-.456.56-.683 1.464-.683 2.712 0 .512.056.984.17 1.416.114.432.285.808.513 1.128.228.32.521.576.88.768.374.176.805.264 1.294.264.537 0 .993-.064 1.368-.192a5.28 5.28 0 001.001-.504l.733 2.376zM61.267 8.5h1.685V6.244l3.175-.888V8.5h2.98v2.64h-2.98v4.608c0 .832.082 1.432.245 1.8.179.352.504.528.977.528.325 0 .602-.032.83-.096.244-.064.513-.16.806-.288l.562 2.4c-.44.208-.953.384-1.539.528a7.446 7.446 0 01-1.783.216c-1.123 0-1.954-.28-2.49-.84-.522-.576-.782-1.512-.782-2.808V11.14h-1.686V8.5zM50.857 9.196c.651-.288 1.424-.512 2.32-.672a14.539 14.539 0 012.809-.264c.846 0 1.555.104 2.124.312.57.192 1.018.472 1.344.84.342.368.578.808.708 1.32.147.512.22 1.088.22 1.728 0 .704-.025 1.416-.073 2.136a45.78 45.78 0 00-.098 2.088c0 .688.024 1.36.073 2.016.049.64.171 1.248.366 1.824h-2.588l-.513-1.656h-.123c-.325.496-.781.928-1.367 1.296-.57.352-1.31.528-2.223.528-.57 0-1.082-.08-1.538-.24a3.59 3.59 0 01-1.173-.72 3.442 3.442 0 01-.757-1.104 3.726 3.726 0 01-.268-1.44c0-.736.162-1.352.488-1.848.342-.512.822-.92 1.441-1.224.635-.32 1.384-.536 2.247-.648.88-.128 1.856-.168 2.93-.12.115-.896.05-1.536-.195-1.92-.244-.4-.79-.6-1.636-.6-.635 0-1.31.064-2.027.192-.7.128-1.278.296-1.734.504l-.757-2.328zm4.03 8.928c.635 0 1.14-.136 1.514-.408.374-.288.651-.592.83-.912v-1.56a9.023 9.023 0 00-1.465-.024 5.357 5.357 0 00-1.221.216c-.359.112-.644.272-.855.48a1.063 1.063 0 00-.318.792c0 .448.13.8.391 1.056.277.24.651.36 1.123.36zM48.586 7.612c0 .416-.057.832-.171 1.248-.098.416-.269.8-.513 1.152a3.95 3.95 0 01-.953.912c-.39.256-.863.448-1.416.576v.144c.488.08.952.216 1.392.408.44.192.822.456 1.148.792.325.336.578.744.757 1.224.195.48.293 1.048.293 1.704 0 .864-.187 1.616-.562 2.256a4.588 4.588 0 01-1.514 1.536c-.619.4-1.327.696-2.125.888-.798.192-1.62.288-2.466.288H41.38c-.423 0-.888-.016-1.392-.048a25.685 25.685 0 01-1.514-.096c-.505-.048-.97-.12-1.393-.216V3.82a27.534 27.534 0 012.345-.264c.456-.032.92-.056 1.392-.072.489-.016.969-.024 1.441-.024.798 0 1.571.064 2.32.192.766.112 1.441.32 2.027.624a3.669 3.669 0 011.441 1.272c.359.544.538 1.232.538 2.064zM42.48 18.028c.407 0 .798-.048 1.172-.144.391-.096.733-.24 1.026-.432.293-.208.53-.464.708-.768.18-.304.269-.664.269-1.08 0-.528-.106-.944-.318-1.248a2.053 2.053 0 00-.83-.696 3.297 3.297 0 00-1.148-.336 11.262 11.262 0 00-1.27-.072h-1.783v4.632c.082.032.212.056.391.072l.562.048c.211 0 .423.008.635.024h.586zm-1.1-7.272c.229 0 .49-.008.782-.024.31-.016.562-.04.758-.072a4.327 4.327 0 001.587-.864c.456-.384.684-.888.684-1.512 0-.416-.082-.76-.244-1.032a1.813 1.813 0 00-.66-.648 2.637 2.637 0 00-.903-.336 5.834 5.834 0 00-1.075-.096c-.423 0-.814.008-1.172.024-.359.016-.635.04-.83.072v4.488h1.074z"
        />
        <defs>
            <linearGradient
                id="paint0_linear_990_5232"
                x1="3.077"
                x2="18.702"
                y1="15.78"
                y2="15.93"
                gradientUnits="userSpaceOnUse"
            >
                <stop stopColor="#00CBEC" />
                <stop offset="0.51" stopColor="#A112FF" />
                <stop offset="1" stopColor="#FF5543" />
            </linearGradient>
            <clipPath id="clip0_990_5232">
                <path fill="#fff" d="M0 0H24.423V24H0z" />
            </clipPath>
        </defs>
    </svg>
)