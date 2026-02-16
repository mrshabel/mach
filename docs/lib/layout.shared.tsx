import type { BaseLayoutProps } from "fumadocs-ui/layouts/shared";

export const gitConfig = {
    user: "mrshabel",
    repo: "mach",
    branch: "main",
};

export function baseOptions(): BaseLayoutProps {
    return {
        nav: {
            title: "Mach",
        },
        githubUrl: `https://github.com/${gitConfig.user}/${gitConfig.repo}`,
    };
}
