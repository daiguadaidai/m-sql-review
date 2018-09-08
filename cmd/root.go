// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/daiguadaidai/m-sql-review/config"
	"github.com/daiguadaidai/m-sql-review/service"
	"github.com/outbrain/golib/log"
)

var runConfig *config.ReviewConfig

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "m-sql-review",
	Short: "SQL审核工具",
	Long: `
    一款SQL审核工具, 主要用于MySQL SQL 相关审核. 启动工具后会提供一个http接口为用户实时链接并且审核相关SQL.
    启动工具:
    ./m-sql-review \
        --rule-name-length=100 \
        --rule-name-reg="^[a-zA-Z\$_][a-zA-Z\$\d_]*$" \
        --rule-charset="utf8,utf8mb4" \
        --rule-collate="utf8_general_ci,utf8mb4_general_ci" \
        --rule-allow-drop-database=false \
        --rule-allow-drop-table=false \
        --rule-allow-rename-table=false \
        --rule-allow-truncate-table=false
    `,
	Run: func(cmd *cobra.Command, args []string) {
		err := service.Run(runConfig)
		if err != nil {
			log.Errorf("运行失败: %v", err)
		}
		fmt.Println(runConfig)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	runConfig = new(config.ReviewConfig)

	rootCmd.Flags().IntVar(&runConfig.RuleNameLength, "rule-name-length",
		config.RULE_NAME_LENGTH,"通用名称长度")
	rootCmd.Flags().StringVar(&runConfig.RuleNameReg, "rule-name-reg",
		config.RULE_NAME_REG, "通用名称匹配规则")
	rootCmd.Flags().StringVar(&runConfig.RuleCharSet, "rule-charset",
		config.RULE_CHARSET,"通用允许的字符集, 默认(多个用逗号隔开)")
	rootCmd.Flags().StringVar(&runConfig.RuleCollate, "rule-collate",
		config.RULE_COLLATE,"通用允许的collate, 默认(多个用逗号隔开)")
	rootCmd.Flags().BoolVar(&runConfig.RuleAllowDropDatabase, "rule-allow-drop-database",
		config.RULE_ALLOW_DROP_DATABASE,
		fmt.Sprintf("是否允许删除数据库, 默认: %v", config.RULE_ALLOW_DROP_DATABASE))
	rootCmd.Flags().BoolVar(&runConfig.RuleAllowDropTable, "rule-allow-drop-table",
		config.RULE_ALLOW_DROP_TABLE,
		fmt.Sprintf("是否允许删除表, 默认: %v", config.RULE_ALLOW_DROP_TABLE))
	rootCmd.Flags().BoolVar(&runConfig.RuleAllowRenameTable, "rule-allow-rename-table",
		config.RULE_ALLOW_RENAME_TABLE,
		fmt.Sprintf("是否允许重命名表, 默认: %v", config.RULE_ALLOW_RENAME_TABLE))
	rootCmd.Flags().BoolVar(&runConfig.RuleAllowTruncateTable, "rule-allow-truncate-table",
		config.RULE_ALLOW_TRUNCATE_TABLE,
		fmt.Sprintf("是否允许truncate表, 默认: %v", config.RULE_ALLOW_TRUNCATE_TABLE))
}
