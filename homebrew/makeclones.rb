require 'rbconfig'
class Makeclones < Formula
  desc ""
  homepage "https://github.com/droxey/makeclones"
  version "1.0.0"

  if Hardware::CPU.is_64_bit?
    case RbConfig::CONFIG['host_os']
    when /mswin|msys|mingw|cygwin|bccwin|wince|emc/
      :windows
    when /darwin|mac os/
      url "https://github.com/droxey/makeclones/releases/download/v1.0.0/makeclones_1.0.0_darwin_amd64.zip"
      sha256 "61ce8cba7656c1bd20e5a14e2e5d979f26643556dc8f89168f1c671838ea4e84"
    when /linux/
      url "https://github.com/droxey/makeclones/releases/download/v1.0.0/"
      sha256 ""
    when /solaris|bsd/
      :unix
    else
      :unknown
    end
  else
    case RbConfig::CONFIG['host_os']
    when /mswin|msys|mingw|cygwin|bccwin|wince|emc/
      :windows
    when /darwin|mac os/
      url "https://github.com/droxey/makeclones/releases/download/v1.0.0/makeclones_1.0.0_darwin_386.zip"
      sha256 "d3be464c4afab53cdacf0e9ce10ca35ddd1d82a42bc355ca3dc64b3a99b23fc8"
    when /linux/
      url "https://github.com/droxey/makeclones/releases/download/v1.0.0/"
      sha256 ""
    when /solaris|bsd/
      :unix
    else
      :unknown
    end
  end

  def install
    bin.install "makeclones"
  end

  test do
    system "makeclones"
  end

end
