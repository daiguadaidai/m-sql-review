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
        --rule-table-engine="innodb"
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
	rootCmd.Flags().StringVar(&runConfig.RuleTableEngine, "rule-table-engine",
		config.RULE_TABLE_ENGINE, "允许的存储引擎 默认(多个用逗号隔开)")
	//不允许的字段用法
	usage := `不允许的字段类型, 填写的是字段类型代码, 如: 不能使用(text,int)使用的配置为: `
	usage += `--rule-not-allow-column-type="252,3". `
	usage += `字段对应代码: Decimal:0, TinyInt:1, ShortInt:2, Int:3, Float:4, Double:5, Null:6, `
	usage += `Timestamp:7, bigint:8, MeduimInt:9, Date:10, Time:11, Datetime:12, Year:13, `
	usage += `NewDate:= 14, Varchar:15, Bit:16, JSON:245, NewDecimal:246, Enum:247, Set:248, `
	usage += `TinyBlob:249, MediumBlob:250, LongBlob:251, Blob:252, VarString:253, String:254, Geometry:255. `
	usage += `注意: blob, text 这些大字段类型代码是一样的`
	rootCmd.Flags().StringVar(&runConfig.RuleTableEngine, "rule-not-allow-column-type",
		config.RULE_NOT_ALLOW_COLUMN_TYPE, usage)
	rootCmd.Flags().BoolVar(&runConfig.RuleNeedTableComment, "rule-need-table-comment",
		config.RULE_NEED_TABLE_COMMENT,
		fmt.Sprintf("表是否需要注释 默认: %v", config.RULE_NEED_TABLE_COMMENT))

}
