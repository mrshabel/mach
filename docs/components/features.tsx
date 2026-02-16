const features = [
    {
        icon: "⚡",
        title: "stdlib routing",
        description:
            "Go 1.22's ServeMux is fast enough. Why reinvent it? Method routing, path params, wildcards—it's all there.",
        color: "blue",
    },
    {
        icon: "♻️",
        title: "context pooling",
        description:
            "sync.Pool reuses Context objects. Less GC pressure, more throughput. Your API stays fast under load.",
        color: "cyan",
    },
    {
        icon: "🔌",
        title: "standard middleware",
        description:
            "Just http.Handler wrapping. No proprietary patterns. Works with any stdlib-compatible middleware.",
        color: "teal",
    },
    {
        icon: "📦",
        title: "zero dependencies",
        description:
            "Nothing to `go get` except Mach. Your dependency tree stays clean. Your builds stay fast.",
        color: "purple",
    },
    {
        icon: "🛡️",
        title: "graceful shutdown",
        description:
            "Built-in signal handling. Your requests finish cleanly when you deploy. No dropped connections.",
        color: "pink",
    },
    {
        icon: "📖",
        title: "learn once, use forever",
        description:
            "No custom abstractions. If you know Go's stdlib, you already know Mach. Onboard in minutes.",
        color: "orange",
    },
];

export function FeaturesSection() {
    return (
        <div className="relative py-32 bg-slate-950">
            <div className="max-w-7xl mx-auto px-6 lg:px-8">
                <div className="text-center mb-16">
                    <h2 className="text-3xl lg:text-5xl font-bold text-white mb-4">
                        Built different
                    </h2>
                    <p className="text-lg text-slate-400">
                        No radix trees, no custom routers, no vendor lock-in
                    </p>
                </div>

                <div className="grid md:grid-cols-2 lg:grid-cols-3 gap-8">
                    {/* Feature 1 */}
                    <div className="group relative bg-gradient-to-br from-slate-900 to-slate-950 p-8 rounded-2xl border border-slate-800 hover:border-blue-500/50 transition-all">
                        <div className="absolute inset-0 bg-gradient-to-br from-blue-500/0 to-blue-500/5 rounded-2xl opacity-0 group-hover:opacity-100 transition-opacity"></div>
                        <div className="relative">
                            <div className="w-12 h-12 bg-blue-500/10 rounded-lg flex items-center justify-center mb-4">
                                <svg
                                    className="w-6 h-6 text-blue-400"
                                    fill="none"
                                    viewBox="0 0 24 24"
                                    stroke="currentColor"
                                >
                                    <path
                                        strokeLinecap="round"
                                        strokeLinejoin="round"
                                        strokeWidth={2}
                                        d="M13 10V3L4 14h7v7l9-11h-7z"
                                    />
                                </svg>
                            </div>
                            <h3 className="text-xl font-semibold text-white mb-3">
                                stdlib routing
                            </h3>
                            <p className="text-slate-400 leading-relaxed">
                                Go 1.22's ServeMux is fast enough. Why reinvent
                                it? Method routing, path params, wildcards—it's
                                all there.
                            </p>
                        </div>
                    </div>

                    {/* Feature 2 */}
                    <div className="group relative bg-gradient-to-br from-slate-900 to-slate-950 p-8 rounded-2xl border border-slate-800 hover:border-cyan-500/50 transition-all">
                        <div className="absolute inset-0 bg-gradient-to-br from-cyan-500/0 to-cyan-500/5 rounded-2xl opacity-0 group-hover:opacity-100 transition-opacity"></div>
                        <div className="relative">
                            <div className="w-12 h-12 bg-cyan-500/10 rounded-lg flex items-center justify-center mb-4">
                                <svg
                                    className="w-6 h-6 text-cyan-400"
                                    fill="none"
                                    viewBox="0 0 24 24"
                                    stroke="currentColor"
                                >
                                    <path
                                        strokeLinecap="round"
                                        strokeLinejoin="round"
                                        strokeWidth={2}
                                        d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"
                                    />
                                </svg>
                            </div>
                            <h3 className="text-xl font-semibold text-white mb-3">
                                context pooling
                            </h3>
                            <p className="text-slate-400 leading-relaxed">
                                sync.Pool reuses Context objects. Less GC
                                pressure, more throughput. Your API stays fast
                                under load.
                            </p>
                        </div>
                    </div>

                    {/* Feature 3 */}
                    <div className="group relative bg-gradient-to-br from-slate-900 to-slate-950 p-8 rounded-2xl border border-slate-800 hover:border-teal-500/50 transition-all">
                        <div className="absolute inset-0 bg-gradient-to-br from-teal-500/0 to-teal-500/5 rounded-2xl opacity-0 group-hover:opacity-100 transition-opacity"></div>
                        <div className="relative">
                            <div className="w-12 h-12 bg-teal-500/10 rounded-lg flex items-center justify-center mb-4">
                                <svg
                                    className="w-6 h-6 text-teal-400"
                                    fill="none"
                                    viewBox="0 0 24 24"
                                    stroke="currentColor"
                                >
                                    <path
                                        strokeLinecap="round"
                                        strokeLinejoin="round"
                                        strokeWidth={2}
                                        d="M12 6V4m0 2a2 2 0 100 4m0-4a2 2 0 110 4m-6 8a2 2 0 100-4m0 4a2 2 0 110-4m0 4v2m0-6V4m6 6v10m6-2a2 2 0 100-4m0 4a2 2 0 110-4m0 4v2m0-6V4"
                                    />
                                </svg>
                            </div>
                            <h3 className="text-xl font-semibold text-white mb-3">
                                standard middleware
                            </h3>
                            <p className="text-slate-400 leading-relaxed">
                                Just http.Handler wrapping. No proprietary
                                patterns. Works with any stdlib-compatible
                                middleware.
                            </p>
                        </div>
                    </div>

                    {/* Feature 4 */}
                    <div className="group relative bg-gradient-to-br from-slate-900 to-slate-950 p-8 rounded-2xl border border-slate-800 hover:border-purple-500/50 transition-all">
                        <div className="absolute inset-0 bg-gradient-to-br from-purple-500/0 to-purple-500/5 rounded-2xl opacity-0 group-hover:opacity-100 transition-opacity"></div>
                        <div className="relative">
                            <div className="w-12 h-12 bg-purple-500/10 rounded-lg flex items-center justify-center mb-4">
                                <svg
                                    className="w-6 h-6 text-purple-400"
                                    fill="none"
                                    viewBox="0 0 24 24"
                                    stroke="currentColor"
                                >
                                    <path
                                        strokeLinecap="round"
                                        strokeLinejoin="round"
                                        strokeWidth={2}
                                        d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4"
                                    />
                                </svg>
                            </div>
                            <h3 className="text-xl font-semibold text-white mb-3">
                                zero dependencies
                            </h3>
                            <p className="text-slate-400 leading-relaxed">
                                Nothing to `go get` except Mach. Your dependency
                                tree stays clean. Your builds stay fast.
                            </p>
                        </div>
                    </div>

                    {/* Feature 5 */}
                    <div className="group relative bg-gradient-to-br from-slate-900 to-slate-950 p-8 rounded-2xl border border-slate-800 hover:border-pink-500/50 transition-all">
                        <div className="absolute inset-0 bg-gradient-to-br from-pink-500/0 to-pink-500/5 rounded-2xl opacity-0 group-hover:opacity-100 transition-opacity"></div>
                        <div className="relative">
                            <div className="w-12 h-12 bg-pink-500/10 rounded-lg flex items-center justify-center mb-4">
                                <svg
                                    className="w-6 h-6 text-pink-400"
                                    fill="none"
                                    viewBox="0 0 24 24"
                                    stroke="currentColor"
                                >
                                    <path
                                        strokeLinecap="round"
                                        strokeLinejoin="round"
                                        strokeWidth={2}
                                        d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z"
                                    />
                                </svg>
                            </div>
                            <h3 className="text-xl font-semibold text-white mb-3">
                                graceful shutdown
                            </h3>
                            <p className="text-slate-400 leading-relaxed">
                                Built-in signal handling. Your requests finish
                                cleanly when you deploy. No dropped connections.
                            </p>
                        </div>
                    </div>

                    {/* Feature 6 */}
                    <div className="group relative bg-gradient-to-br from-slate-900 to-slate-950 p-8 rounded-2xl border border-slate-800 hover:border-orange-500/50 transition-all">
                        <div className="absolute inset-0 bg-gradient-to-br from-orange-500/0 to-orange-500/5 rounded-2xl opacity-0 group-hover:opacity-100 transition-opacity"></div>
                        <div className="relative">
                            <div className="w-12 h-12 bg-orange-500/10 rounded-lg flex items-center justify-center mb-4">
                                <svg
                                    className="w-6 h-6 text-orange-400"
                                    fill="none"
                                    viewBox="0 0 24 24"
                                    stroke="currentColor"
                                >
                                    <path
                                        strokeLinecap="round"
                                        strokeLinejoin="round"
                                        strokeWidth={2}
                                        d="M12 6.253v13m0-13C10.832 5.477 9.246 5 7.5 5S4.168 5.477 3 6.253v13C4.168 18.477 5.754 18 7.5 18s3.332.477 4.5 1.253m0-13C13.168 5.477 14.754 5 16.5 5c1.747 0 3.332.477 4.5 1.253v13C19.832 18.477 18.247 18 16.5 18c-1.746 0-3.332.477-4.5 1.253"
                                    />
                                </svg>
                            </div>
                            <h3 className="text-xl font-semibold text-white mb-3">
                                learn once, use forever
                            </h3>
                            <p className="text-slate-400 leading-relaxed">
                                No custom abstractions. If you know Go's stdlib,
                                you already know Mach. Onboard in minutes.
                            </p>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    );
}
